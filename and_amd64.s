// Code generated by command: go run src.go -out ../and_amd64.s -stubs ../and_stubs_amd64.go -pkg and. DO NOT EDIT.

#include "textflag.h"

// func andAVX2(dst *byte, a *byte, b *byte, l uint64)
// Requires: AVX, AVX2
TEXT ·andAVX2(SB), NOSPLIT, $0-32
	MOVQ a+8(FP), AX
	MOVQ b+16(FP), CX
	MOVQ dst+0(FP), DX
	MOVQ l+24(FP), BX

loop:
	VMOVDQU (AX), Y0
	VMOVDQU (CX), Y8
	VMOVDQU 32(AX), Y1
	VMOVDQU 32(CX), Y9
	VMOVDQU 64(AX), Y2
	VMOVDQU 64(CX), Y10
	VMOVDQU 96(AX), Y3
	VMOVDQU 96(CX), Y11
	VMOVDQU 128(AX), Y4
	VMOVDQU 128(CX), Y12
	VMOVDQU 160(AX), Y5
	VMOVDQU 160(CX), Y13
	VMOVDQU 192(AX), Y6
	VMOVDQU 192(CX), Y14
	VMOVDQU 224(AX), Y7
	VMOVDQU 224(CX), Y15
	VPAND   Y0, Y8, Y8
	VPAND   Y1, Y9, Y9
	VPAND   Y2, Y10, Y10
	VPAND   Y3, Y11, Y11
	VPAND   Y4, Y12, Y12
	VPAND   Y5, Y13, Y13
	VPAND   Y6, Y14, Y14
	VPAND   Y7, Y15, Y15
	VMOVDQU (DX), Y8
	VMOVDQU 32(DX), Y9
	VMOVDQU 64(DX), Y10
	VMOVDQU 96(DX), Y11
	VMOVDQU 128(DX), Y12
	VMOVDQU 160(DX), Y13
	VMOVDQU 192(DX), Y14
	VMOVDQU 224(DX), Y15
	SUBQ    $0x00000001, BX
	JNZ     loop
	RET
