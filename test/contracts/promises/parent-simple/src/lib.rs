#![no_std]
#![no_main]
#![allow(unused_attributes)]
#![feature(lang_items)]

use elrond_wasm::api::{EndpointArgumentApi, EndpointFinishApi, StorageReadApi, StorageWriteApi};
use elrond_wasm_node::ArwenApiImpl;

use promises_common::*;

pub static EEI: ArwenApiImpl = ArwenApiImpl{};

const SUCCESS_CALLBACK_ARGUMENT_KEY: &[u8] = b"SuccessCallbackArg";
const FAIL_CALLBACK_ARGUMENT_KEY: &[u8] = b"FailCallbackArg";
const CURRENT_STORAGE_INDEX_KEY: &[u8] = b"CurrentStorageIndex";

#[no_mangle]
pub extern "C" fn no_async() {
    EEI.finish_i64(42);
}

// one async call

#[no_mangle]
pub extern "C" fn one_async_call_no_cb_with_call_value() {
    let mut value = [0u8; 32];
    value[31] = 16;

    create_async_call("testgroup",
                      &Address::from(CHILD_ADDRESS),
                      &value,
                      b"answer",
                      "",
                      "",
                      100000);
}

#[no_mangle]
pub extern "C" fn one_async_call_no_cb_fail() {
    create_async_call("testgroup",
                      &Address::from(CHILD_ADDRESS),
                      &ZERO,
                      b"fail",
                      "",
                      "",
                      100000);
}

#[no_mangle]
pub extern "C" fn one_async_call_no_cb_fail_with_call_value() {
    let mut value = [0u8; 32];
    value[31] = 16;

    create_async_call("testgroup",
                      &Address::from(CHILD_ADDRESS),
                      &value,
                      b"fail",
                      "",
                      "",
                      100000);
}

#[no_mangle]
pub extern "C" fn one_async_call_success_cb() {
    create_async_call("testgroup",
                      &Address::from(CHILD_ADDRESS),
                      &ZERO,
                      b"answer",
                      "success_callback_one_arg",
                      "",
                      100000);
}

#[no_mangle]
pub extern "C" fn one_async_call_fail_cb() {
    create_async_call("testgroup",
                      &Address::from(CHILD_ADDRESS),
                      &ZERO,
                      b"answer",
                      "",
                      "fail_callback",
                      100000);
}

// two async calls

pub extern "C" fn two_async_same_cb_success_both() {
    create_async_call("testgroup",
                      &Address::from(CHILD_ADDRESS),
                      &ZERO,
                      b"echo@0x01",
                      "success_callback_one_arg",
                      "fail_callback",
                      100000);
    
    create_async_call("testgroup",
                      &Address::from(CHILD_ADDRESS),
                      &ZERO,
                      b"echo@0x02",
                      "success_callback_one_arg",
                      "fail_callback",
                      100000);
}

pub extern "C" fn two_async_same_cb_success_first_fail_second() {
    create_async_call("testgroup",
                      &Address::from(CHILD_ADDRESS),
                      &ZERO,
                      b"echo@0x01",
                      "success_callback_one_arg",
                      "fail_callback",
                      100000);
    
    create_async_call("testgroup",
                      &Address::from(CHILD_ADDRESS),
                      &ZERO,
                      b"fail",
                      "success_callback_one_arg",
                      "fail_callback",
                      100000);
}

pub extern "C" fn two_async_same_cb_fail_first_success_second() {
    create_async_call("testgroup",
                      &Address::from(CHILD_ADDRESS),
                      &ZERO,
                      b"fail",
                      "success_callback_one_arg",
                      "fail_callback",
                      100000);
    
    create_async_call("testgroup",
                      &Address::from(CHILD_ADDRESS),
                      &ZERO,
                      b"echo@0x02",
                      "success_callback_one_arg",
                      "fail_callback",
                      100000);
}

pub extern "C" fn two_async_same_cb_fail_both() {
    create_async_call("testgroup",
                      &Address::from(CHILD_ADDRESS),
                      &ZERO,
                      b"fail",
                      "success_callback_one_arg",
                      "fail_callback",
                      100000);
    
    create_async_call("testgroup",
                      &Address::from(CHILD_ADDRESS),
                      &ZERO,
                      b"fail",
                      "success_callback_one_arg",
                      "fail_callback",
                      100000);
}

// callbacks

// first argument is "0" for success, followed by data passed by finish() in callee contract
#[no_mangle]
pub extern "C" fn success_callback_one_arg() {
    let expected_num_args = 2;
    EEI.check_num_arguments(expected_num_args);

    let mut storage_index = EEI.storage_load_u64(&CURRENT_STORAGE_INDEX_KEY);

    for arg_index in 0..expected_num_args {
        let arg = EEI.get_argument_u64(arg_index);
        let storage_key = construct_storage_key(&[SUCCESS_CALLBACK_ARGUMENT_KEY, &[storage_index as u8]]);

        storage_index += 1;
        EEI.storage_store_u64(&storage_key, arg);
    }

    EEI.storage_store_u64(&CURRENT_STORAGE_INDEX_KEY, storage_index);
}

// first argument is error code, followed by error message
#[no_mangle]
pub extern "C" fn fail_callback() {
    let expected_num_args = 2;
    EEI.check_num_arguments(expected_num_args);

    let mut storage_index = EEI.storage_load_u64(&CURRENT_STORAGE_INDEX_KEY);

    for arg_index in 0..expected_num_args {
        let arg = EEI.get_argument_vec_u8(arg_index);
        let storage_key = construct_storage_key(&[FAIL_CALLBACK_ARGUMENT_KEY, &[storage_index as u8]]);
    
        storage_index += 1;
        EEI.storage_store_slice_u8(&storage_key, &arg);
    }

    EEI.storage_store_u64(&CURRENT_STORAGE_INDEX_KEY, storage_index);
}

