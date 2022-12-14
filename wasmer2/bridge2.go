package wasmer2

// #cgo LDFLAGS: -Wl,-rpath,${SRCDIR} -L${SRCDIR}
// #cgo linux,amd64 LDFLAGS:-lvmexeccapi
// #cgo darwin,amd64 LDFLAGS:-lvmexeccapi
// #include "./libvmexeccapi.h"
//
import "C"
import (
	"unsafe"
)

type cBool C.bool
type cChar C.char
type cInt C.int
type cUchar C.uchar
type cUint C.uint
type cUint32T C.uint32_t
type cUint8T C.uint8_t

type cWasmerExecutorT C.vm_exec_executor_t
type cWasmerInstanceT C.vm_exec_instance_t
type cWasmerOpcodeCostT C.vm_exec_opcode_cost_t
type cWasmerResultT C.vm_exec_result_t
type cWasmerCompilationOptions C.vm_exec_compilation_options_t
type cWasmerVmHookPointers = C.vm_exec_vm_hook_c_func_pointers

// type cWasmerByteArray C.vm_exec_byte_array
// type cWasmerInstanceContextT C.vm_exec_instance_context_t
// type cWasmerMemoryT C.vm_exec_memory_t
// type cFuncGetGasLeft = C.get_gas_left_func

const cWasmerOk = C.VM_EXEC_OK

func cWasmerExecutorSetOpcodeCost(
	executor *cWasmerExecutorT,
	opcodeCost *cWasmerOpcodeCostT,
) cWasmerResultT {
	return (cWasmerResultT)(C.vm_exec_set_opcode_costs(
		(*C.vm_exec_executor_t)(executor),
		(*C.vm_exec_opcode_cost_t)(opcodeCost),
	))
}

func cWasmerInstanceSetGasLimit(instance *cWasmerInstanceT, gasLimit uint64) cWasmerResultT {
	return (cWasmerResultT)(C.vm_exec_instance_set_points_limit(
		(*C.vm_exec_instance_t)(instance),
		(C.uint64_t)(gasLimit),
	))
}

func cWasmerInstanceSetPointsUsed(instance *cWasmerInstanceT, points uint64) cWasmerResultT {
	return (cWasmerResultT)(C.vm_exec_instance_set_points_used(
		(*C.vm_exec_instance_t)(instance),
		(C.uint64_t)(points),
	))
}

func cWasmerInstanceGetPointsUsed(instance *cWasmerInstanceT) uint64 {
	return uint64(C.vm_exec_instance_get_points_used(
		(*C.vm_exec_instance_t)(instance),
	))
}

func cWasmerInstanceSetBreakpointValue(instance *cWasmerInstanceT, value uint64) cWasmerResultT {
	return (cWasmerResultT)(C.vm_exec_instance_set_breakpoint_value(
		(*C.vm_exec_instance_t)(instance),
		(C.uint64_t)(value),
	))
}

func cWasmerInstanceGetBreakpointValue(instance *cWasmerInstanceT) uint64 {
	return uint64(C.vm_exec_instance_get_breakpoint_value(
		(*C.vm_exec_instance_t)(instance),
	))
}

func cWasmerInstanceCache(
	instance *cWasmerInstanceT,
	cacheBytes **cUchar,
	cacheLen *cUint32T,
) cWasmerResultT {
	return (cWasmerResultT)(C.vm_exec_instance_cache(
		(*C.vm_exec_instance_t)(instance),
		(**C.uchar)(unsafe.Pointer(cacheBytes)),
		(*C.uint32_t)(cacheLen),
	))
}

func cWasmerInstanceFromCache(
	executor *cWasmerExecutorT,
	instance **cWasmerInstanceT,
	cacheBytes *cUchar,
	cacheLen cUint32T,
	options *cWasmerCompilationOptions,
) cWasmerResultT {
	return (cWasmerResultT)(C.vm_exec_instance_from_cache(
		(*C.vm_exec_executor_t)(executor),
		(**C.vm_exec_instance_t)(unsafe.Pointer(instance)),
		(*C.uchar)(cacheBytes),
		(C.uint32_t)(cacheLen),
		(*C.vm_exec_compilation_options_t)(options),
	))
}

func cWasmerNewExecutor(
	executor **cWasmerExecutorT,
	vmHookPointersPtrPtr unsafe.Pointer,
) cWasmerResultT {
	return (cWasmerResultT)(C.vm_exec_new_executor(
		(**C.vm_exec_executor_t)(unsafe.Pointer(executor)),
		(**C.vm_exec_vm_hook_c_func_pointers)(vmHookPointersPtrPtr),
	))
}

func cWasmerInstantiateWithOptions(
	executor *cWasmerExecutorT,
	instance **cWasmerInstanceT,
	wasmBytes *cUchar,
	wasmBytesLength cUint,
	options *cWasmerCompilationOptions,
) cWasmerResultT {
	return (cWasmerResultT)(C.vm_exec_new_instance(
		(*C.vm_exec_executor_t)(executor),
		(**C.vm_exec_instance_t)(unsafe.Pointer(instance)),
		(*C.uchar)(wasmBytes),
		(C.uint)(wasmBytesLength),
		(*C.vm_exec_compilation_options_t)(options),
	))
}

func cWasmerInstanceCall(
	instance *cWasmerInstanceT,
	name *cChar,
) cWasmerResultT {
	return (cWasmerResultT)(C.vm_exec_instance_call(
		(*C.vm_exec_instance_t)(instance),
		(*C.char)(name),
	))
}

func cWasmerInstanceHasFunction(
	instance *cWasmerInstanceT,
	name *cChar,
) cInt {
	return (cInt)(C.vm_exec_instance_has_function(
		(*C.vm_exec_instance_t)(instance),
		(*C.char)(name),
	))
}

func cWasmerInstanceExportedFunctionNamesLength(instance *cWasmerInstanceT) cInt {
	return (cInt)(C.vm_exported_function_names_length(
		(*C.vm_exec_instance_t)(instance),
	))
}

func cWasmerInstanceExportedFunctionNames(instance *cWasmerInstanceT, buffer *cChar, length cInt) cInt {
	return (cInt)(C.vm_exported_function_names(
		(*C.vm_exec_instance_t)(instance),
		(*C.char)(buffer),
		(C.int)(length),
	))
}

func cWasmerCheckSignatures(instance *cWasmerInstanceT) cInt {
	return (cInt)(C.vm_check_signatures(
		(*C.vm_exec_instance_t)(instance),
	))
}

func cWasmerExecutorContextDataSet(executor *cWasmerExecutorT, vmHooksPtr unsafe.Pointer) {
	C.vm_exec_executor_set_vm_hooks_ptr(
		(*C.vm_exec_executor_t)(executor),
		vmHooksPtr,
	)
}

func cWasmerInstanceDestroy(instance *cWasmerInstanceT) {
	C.vm_exec_instance_destroy(
		(*C.vm_exec_instance_t)(instance),
	)
}

func cWasmerInstanceReset(instance *cWasmerInstanceT) cWasmerResultT {
	return (cWasmerResultT)(C.vm_exec_instance_reset(
		(*C.vm_exec_instance_t)(instance),
	))
}

func cWasmerLastErrorLength() cInt {
	return (cInt)(C.vm_exec_last_error_length())
}

func cWasmerLastErrorMessage(buffer *cChar, length cInt) cInt {
	return (cInt)(C.vm_exec_last_error_message(
		(*C.char)(buffer),
		(C.int)(length),
	))
}

func cWasmerMemoryDataLength(instance *cWasmerInstanceT) cUint32T {
	return (cUint32T)(C.vm_exec_instance_memory_data_length(
		(*C.vm_exec_instance_t)(instance),
	))
}

func cWasmerMemoryData(instance *cWasmerInstanceT) *cUint8T {
	return (*cUint8T)(C.vm_exec_instance_memory_data(
		(*C.vm_exec_instance_t)(instance),
	))
}

func cWasmerMemoryGrow(instance *cWasmerInstanceT, numberOfPages cUint32T) cWasmerResultT {
	return (cWasmerResultT)(C.vm_exec_instance_memory_grow(
		(*C.vm_exec_instance_t)(instance),
		(C.uint32_t)(numberOfPages),
	))
}

// func cWasmerInstanceIsFunctionImported(instance *cWasmerInstanceT, name string) bool {
// 	var functionName = cCString(name)
// 	return bool(C.vm_exec_instance_is_function_imported(
// 		(*C.vm_exec_instance_t)(instance),
// 		(*C.char)(unsafe.Pointer(functionName)),
// 	))
// }

// func cWasmerExportToMemory(export *cWasmerExportT, memory **cWasmerMemoryT) cWasmerResultT {
// 	return (cWasmerResultT)(C.vm_exec_export_to_memory(
// 		(*C.vm_exec_export_t)(export),
// 		(**C.vm_exec_memory_t)(unsafe.Pointer(memory)),
// 	))
// }

// func cWasmerMemoryDestroy(memory *cWasmerMemoryT) {
// 	C.vm_exec_memory_destroy(
// 		(*C.vm_exec_memory_t)(memory),
// 	)
// }

// func cWasmerInstanceContextGet(instance *cWasmerInstanceT) *cWasmerInstanceContextT {
// 	return (*cWasmerInstanceContextT)(C.vm_exec_instance_context_get(
// 		(*C.vm_exec_instance_t)(instance),
// 	))
// }

// func cWasmerInstanceContextDataGet(instanceContext *cWasmerInstanceContextT) unsafe.Pointer {
// 	return unsafe.Pointer(C.vm_exec_instance_context_data_get(
// 		(*C.vm_exec_instance_context_t)(instanceContext),
// 	))
// }

// func cWasmerInstanceContextMemory(instanceContext *cWasmerInstanceContextT) *cWasmerMemoryT {
// 	return (*cWasmerMemoryT)(C.vm_exec_instance_context_memory(
// 		(*C.vm_exec_instance_context_t)(instanceContext),
// 		0,
// 	))
// }

func cCString(string string) *cChar {
	return (*cChar)(C.CString(string))
}

func cFree(pointer unsafe.Pointer) {
	C.free(pointer)
}

func cGoString(string *cChar) string {
	return C.GoString((*C.char)(string))
}

func cGoStringN(string *cChar, length cInt) string {
	return C.GoStringN((*C.char)(string), (C.int)(length))
}

// func cGoStringToWasmerByteArray(string string) cWasmerByteArray {
// 	var cString = cCString(string)

// 	var byteArray cWasmerByteArray
// 	byteArray.bytes = (*C.uchar)(unsafe.Pointer(cString))
// 	byteArray.bytes_len = (C.uint)(len(string))

// 	return byteArray
// }