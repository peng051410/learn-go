TEXT runtime.main.func2(SB) /usr/local/opt/go/libexec/src/runtime/proc.go
  proc.go:203		0x1032880		493b6610		CMPQ 0x10(R14), SP		
  proc.go:203		0x1032884		7628			JBE 0x10328ae			
  proc.go:203		0x1032886		4883ec08		SUBQ $0x8, SP			
  proc.go:203		0x103288a		48892c24		MOVQ BP, 0(SP)			
  proc.go:203		0x103288e		488d2c24		LEAQ 0(SP), BP			
  proc.go:203		0x1032892		488b4208		MOVQ 0x8(DX), AX		
  proc.go:204		0x1032896		803800			CMPB $0x0, 0(AX)		
  proc.go:204		0x1032899		740a			JE 0x10328a5			
  proc.go:205		0x103289b		0f1f440000		NOPL 0(AX)(AX*1)		
  proc.go:205		0x10328a0		e8bb940000		CALL runtime.unlockOSThread(SB)	
  proc.go:207		0x10328a5		488b2c24		MOVQ 0(SP), BP			
  proc.go:207		0x10328a9		4883c408		ADDQ $0x8, SP			
  proc.go:207		0x10328ad		c3			RET				
  proc.go:203		0x10328ae		e84d920200		CALL runtime.morestack.abi0(SB)	
  proc.go:203		0x10328b3		ebcb			JMP runtime.main.func2(SB)	
  :-1			0x10328b5		cc			INT $0x3			
  :-1			0x10328b6		cc			INT $0x3			
  :-1			0x10328b7		cc			INT $0x3			
  :-1			0x10328b8		cc			INT $0x3			
  :-1			0x10328b9		cc			INT $0x3			
  :-1			0x10328ba		cc			INT $0x3			
  :-1			0x10328bb		cc			INT $0x3			
  :-1			0x10328bc		cc			INT $0x3			
  :-1			0x10328bd		cc			INT $0x3			
  :-1			0x10328be		cc			INT $0x3			
  :-1			0x10328bf		cc			INT $0x3			

TEXT runtime.main.func1(SB) /usr/local/opt/go/libexec/src/runtime/proc.go
  proc.go:170		0x10580c0		493b6610		CMPQ 0x10(R14), SP			
  proc.go:170		0x10580c4		762d			JBE 0x10580f3				
  proc.go:170		0x10580c6		4883ec20		SUBQ $0x20, SP				
  proc.go:170		0x10580ca		48896c2418		MOVQ BP, 0x18(SP)			
  proc.go:170		0x10580cf		488d6c2418		LEAQ 0x18(SP), BP			
  proc.go:171		0x10580d4		488d05b57e0500		LEAQ go:func.*+1304(SB), AX		
  proc.go:171		0x10580db		31db			XORL BX, BX				
  proc.go:171		0x10580dd		48c7c1ffffffff		MOVQ $-0x1, CX				
  proc.go:171		0x10580e4		e8b7e3fdff		CALL runtime.newm(SB)			
  proc.go:172		0x10580e9		488b6c2418		MOVQ 0x18(SP), BP			
  proc.go:172		0x10580ee		4883c420		ADDQ $0x20, SP				
  proc.go:172		0x10580f2		c3			RET					
  proc.go:170		0x10580f3		e8a83a0000		CALL runtime.morestack_noctxt.abi0(SB)	
  proc.go:170		0x10580f8		ebc6			JMP runtime.main.func1(SB)		
  :-1			0x10580fa		cc			INT $0x3				
  :-1			0x10580fb		cc			INT $0x3				
  :-1			0x10580fc		cc			INT $0x3				
  :-1			0x10580fd		cc			INT $0x3				
  :-1			0x10580fe		cc			INT $0x3				
  :-1			0x10580ff		cc			INT $0x3				

TEXT main.main(SB) /Users/tomyli/github/learn-go/helloworld/main.go
  main.go:5		0x108ddc0		4c8d6424d0			LEAQ -0x30(SP), R12				
  main.go:5		0x108ddc5		4d3b6610			CMPQ 0x10(R14), R12				
  main.go:5		0x108ddc9		0f8646010000			JBE 0x108df15					
  main.go:5		0x108ddcf		4881ecb0000000			SUBQ $0xb0, SP					
  main.go:5		0x108ddd6		4889ac24a8000000		MOVQ BP, 0xa8(SP)				
  main.go:5		0x108ddde		488dac24a8000000		LEAQ 0xa8(SP), BP				
  main.go:6		0x108dde6		48c744241800000000		MOVQ $0x0, 0x18(SP)				
  main.go:7		0x108ddef		440f117c2448			MOVUPS X15, 0x48(SP)				
  main.go:7		0x108ddf5		440f117c2458			MOVUPS X15, 0x58(SP)				
  main.go:7		0x108ddfb		488d442458			LEAQ 0x58(SP), AX				
  main.go:7		0x108de00		4889442438			MOVQ AX, 0x38(SP)				
  main.go:7		0x108de05		8400				TESTB AL, 0(AX)					
  main.go:7		0x108de07		488d1552800000			LEAQ runtime.rodata+32256(SB), DX		
  main.go:7		0x108de0e		4889542458			MOVQ DX, 0x58(SP)				
  main.go:7		0x108de13		488d1586840300			LEAQ runtime.buildVersion.str+112(SB), DX	
  main.go:7		0x108de1a		4889542460			MOVQ DX, 0x60(SP)				
  main.go:7		0x108de1f		8400				TESTB AL, 0(AX)					
  main.go:7		0x108de21		eb00				JMP 0x108de23					
  main.go:7		0x108de23		4889842490000000		MOVQ AX, 0x90(SP)				
  main.go:7		0x108de2b		48c784249800000001000000	MOVQ $0x1, 0x98(SP)				
  main.go:7		0x108de37		48c78424a000000001000000	MOVQ $0x1, 0xa0(SP)				
  main.go:7		0x108de43		bb01000000			MOVL $0x1, BX					
  main.go:7		0x108de48		4889d9				MOVQ BX, CX					
  main.go:7		0x108de4b		e870acffff			CALL fmt.Println(SB)				
  main.go:7		0x108de50		4889442420			MOVQ AX, 0x20(SP)				
  main.go:7		0x108de55		48895c2448			MOVQ BX, 0x48(SP)				
  main.go:7		0x108de5a		48894c2450			MOVQ CX, 0x50(SP)				
  main.go:7		0x108de5f		488b542420			MOVQ 0x20(SP), DX				
  main.go:7		0x108de64		4889542428			MOVQ DX, 0x28(SP)				
  main.go:7		0x108de69		488b542448			MOVQ 0x48(SP), DX				
  main.go:7		0x108de6e		4889542468			MOVQ DX, 0x68(SP)				
  main.go:7		0x108de73		48894c2470			MOVQ CX, 0x70(SP)				
  main.go:7		0x108de78		488b542428			MOVQ 0x28(SP), DX				
  main.go:7		0x108de7d		4889542418			MOVQ DX, 0x18(SP)				
  main.go:8		0x108de82		440f117c2458			MOVUPS X15, 0x58(SP)				
  main.go:8		0x108de88		488d542458			LEAQ 0x58(SP), DX				
  main.go:8		0x108de8d		4889542430			MOVQ DX, 0x30(SP)				
  main.go:8		0x108de92		488b442418			MOVQ 0x18(SP), AX				
  main.go:8		0x108de97		e844b8f7ff			CALL runtime.convT64(SB)			
  main.go:8		0x108de9c		4889442440			MOVQ AX, 0x40(SP)				
  main.go:8		0x108dea1		488b7c2430			MOVQ 0x30(SP), DI				
  main.go:8		0x108dea6		8407				TESTB AL, 0(DI)					
  main.go:8		0x108dea8		488d1571790000			LEAQ runtime.rodata+30656(SB), DX		
  main.go:8		0x108deaf		488917				MOVQ DX, 0(DI)					
  main.go:8		0x108deb2		488d5708			LEAQ 0x8(DI), DX				
  main.go:8		0x108deb6		833d93310e0000			CMPL $0x0, runtime.writeBarrier(SB)		
  main.go:8		0x108debd		7403				JE 0x108dec2					
  main.go:8		0x108debf		90				NOPL						
  main.go:8		0x108dec0		eb06				JMP 0x108dec8					
  main.go:8		0x108dec2		48894708			MOVQ AX, 0x8(DI)				
  main.go:8		0x108dec6		eb0a				JMP 0x108ded2					
  main.go:8		0x108dec8		4889d7				MOVQ DX, DI					
  main.go:8		0x108decb		e8b0fcfcff			CALL runtime.gcWriteBarrier(SB)			
  main.go:8		0x108ded0		eb00				JMP 0x108ded2					
  main.go:8		0x108ded2		488b442430			MOVQ 0x30(SP), AX				
  main.go:8		0x108ded7		8400				TESTB AL, 0(AX)					
  main.go:8		0x108ded9		eb00				JMP 0x108dedb					
  main.go:8		0x108dedb		4889442478			MOVQ AX, 0x78(SP)				
  main.go:8		0x108dee0		48c784248000000001000000	MOVQ $0x1, 0x80(SP)				
  main.go:8		0x108deec		48c784248800000001000000	MOVQ $0x1, 0x88(SP)				
  main.go:8		0x108def8		bb01000000			MOVL $0x1, BX					
  main.go:8		0x108defd		4889d9				MOVQ BX, CX					
  main.go:8		0x108df00		e8bbabffff			CALL fmt.Println(SB)				
  main.go:9		0x108df05		488bac24a8000000		MOVQ 0xa8(SP), BP				
  main.go:9		0x108df0d		4881c4b0000000			ADDQ $0xb0, SP					
  main.go:9		0x108df14		c3				RET						
  main.go:5		0x108df15		e886dcfcff			CALL runtime.morestack_noctxt.abi0(SB)		
  main.go:5		0x108df1a		e9a1feffff			JMP main.main(SB)				
