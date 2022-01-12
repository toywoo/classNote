	.text
	.def	 @feat.00;
	.scl	3;
	.type	0;
	.endef
	.globl	@feat.00
.set @feat.00, 1
	.file	"main.c"
	.def	 _add;
	.scl	2;
	.type	32;
	.endef
	.globl	_add                            # -- Begin function add
	.p2align	4, 0x90
_add:                                   # @add
# %bb.0:
	pushl	%ebp
	movl	%esp, %ebp
	movl	12(%ebp), %eax
	movl	8(%ebp), %eax
	movl	8(%ebp), %eax
	addl	12(%ebp), %eax
	popl	%ebp
	retl
                                        # -- End function
	.def	 _main;
	.scl	2;
	.type	32;
	.endef
	.globl	_main                           # -- Begin function main
	.p2align	4, 0x90
_main:                                  # @main
# %bb.0:
	pushl	%ebp
	movl	%esp, %ebp
	subl	$20, %esp
	movl	$0, -4(%ebp)
	movl	$1, -8(%ebp)
	movl	$2, -12(%ebp)
	movl	-12(%ebp), %eax
	movl	-8(%ebp), %ecx
	movl	%ecx, (%esp)
	movl	%eax, 4(%esp)
	calll	_add
	movl	%eax, -8(%ebp)
	xorl	%eax, %eax
	addl	$20, %esp
	popl	%ebp
	retl
                                        # -- End function
	.addrsig
	.addrsig_sym _add
