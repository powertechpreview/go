// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ppc64

import (
	"cmd/compile/internal/gc"
	"cmd/internal/obj"
	"cmd/internal/obj/ppc64"
)

const (
	LeftRdwr  uint32 = gc.LeftRead | gc.LeftWrite
	RightRdwr uint32 = gc.RightRead | gc.RightWrite
)

// This table gives the basic information about instruction
// generated by the compiler and processed in the optimizer.
// See opt.h for bit definitions.
//
// Instructions not generated need not be listed.
// As an exception to that rule, we typically write down all the
// size variants of an operation even if we just use a subset.
//
// The table is formatted for 8-space tabs.
var progtable = [ppc64.ALAST & obj.AMask]gc.ProgInfo{
	obj.ATYPE:     {Flags: gc.Pseudo | gc.Skip},
	obj.ATEXT:     {Flags: gc.Pseudo},
	obj.AFUNCDATA: {Flags: gc.Pseudo},
	obj.APCDATA:   {Flags: gc.Pseudo},
	obj.AUNDEF:    {Flags: gc.Break},
	obj.AUSEFIELD: {Flags: gc.OK},
	obj.AVARDEF:   {Flags: gc.Pseudo | gc.RightWrite},
	obj.AVARKILL:  {Flags: gc.Pseudo | gc.RightWrite},
	obj.AVARLIVE:  {Flags: gc.Pseudo | gc.LeftRead},

	// NOP is an internal no-op that also stands
	// for USED and SET annotations, not the Power opcode.
	obj.ANOP: {Flags: gc.LeftRead | gc.RightWrite},

	// Integer
	ppc64.AADD & obj.AMask:    {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.AADDC & obj.AMask:   {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.ASUB & obj.AMask:    {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.AADDME & obj.AMask:  {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.ANEG & obj.AMask:    {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.AAND & obj.AMask:    {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.AANDN & obj.AMask:   {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.AOR & obj.AMask:     {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.AORN & obj.AMask:    {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.AXOR & obj.AMask:    {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.AEQV & obj.AMask:    {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.AMULLD & obj.AMask:  {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.AMULLW & obj.AMask:  {Flags: gc.SizeL | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.AMULHD & obj.AMask:  {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.AMULHDU & obj.AMask: {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.AMULHW & obj.AMask:  {Flags: gc.SizeL | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.AMULHWU & obj.AMask: {Flags: gc.SizeL | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.ADIVD & obj.AMask:   {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.ADIVDU & obj.AMask:  {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.ADIVW & obj.AMask:   {Flags: gc.SizeL | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.ADIVWU & obj.AMask:  {Flags: gc.SizeL | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.ASLD & obj.AMask:    {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.ASRD & obj.AMask:    {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.ASRAD & obj.AMask:   {Flags: gc.SizeQ | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.ASLW & obj.AMask:    {Flags: gc.SizeL | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.ASRW & obj.AMask:    {Flags: gc.SizeL | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.ASRAW & obj.AMask:   {Flags: gc.SizeL | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.ACMP & obj.AMask:    {Flags: gc.SizeQ | gc.LeftRead | gc.RightRead},
	ppc64.ACMPU & obj.AMask:   {Flags: gc.SizeQ | gc.LeftRead | gc.RightRead},
	ppc64.ACMPW & obj.AMask:   {Flags: gc.SizeL | gc.LeftRead | gc.RightRead},
	ppc64.ACMPWU & obj.AMask:  {Flags: gc.SizeL | gc.LeftRead | gc.RightRead},
	ppc64.ATD & obj.AMask:     {Flags: gc.SizeQ | gc.RightRead},

	// Floating point.
	ppc64.AFADD & obj.AMask:   {Flags: gc.SizeD | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.AFADDS & obj.AMask:  {Flags: gc.SizeF | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.AFSUB & obj.AMask:   {Flags: gc.SizeD | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.AFSUBS & obj.AMask:  {Flags: gc.SizeF | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.AFMUL & obj.AMask:   {Flags: gc.SizeD | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.AFMULS & obj.AMask:  {Flags: gc.SizeF | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.AFDIV & obj.AMask:   {Flags: gc.SizeD | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.AFDIVS & obj.AMask:  {Flags: gc.SizeF | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.AFCTIDZ & obj.AMask: {Flags: gc.SizeF | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.AFCTIWZ & obj.AMask: {Flags: gc.SizeF | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.AFCFID & obj.AMask:  {Flags: gc.SizeF | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.AFCFIDU & obj.AMask: {Flags: gc.SizeF | gc.LeftRead | gc.RegRead | gc.RightWrite},
	ppc64.AFCMPU & obj.AMask:  {Flags: gc.SizeD | gc.LeftRead | gc.RightRead},
	ppc64.AFRSP & obj.AMask:   {Flags: gc.SizeD | gc.LeftRead | gc.RightWrite | gc.Conv},
	ppc64.AFSQRT & obj.AMask:  {Flags: gc.SizeD | gc.LeftRead | gc.RightWrite},
	ppc64.AFNEG & obj.AMask:   {Flags: gc.SizeD | gc.LeftRead | gc.RightWrite},

	// Moves
	ppc64.AMOVB & obj.AMask:  {Flags: gc.SizeB | gc.LeftRead | gc.RightWrite | gc.Move | gc.Conv},
	ppc64.AMOVBU & obj.AMask: {Flags: gc.SizeB | gc.LeftRead | gc.RightWrite | gc.Move | gc.Conv | gc.PostInc},
	ppc64.AMOVBZ & obj.AMask: {Flags: gc.SizeB | gc.LeftRead | gc.RightWrite | gc.Move | gc.Conv},
	ppc64.AMOVH & obj.AMask:  {Flags: gc.SizeW | gc.LeftRead | gc.RightWrite | gc.Move | gc.Conv},
	ppc64.AMOVHU & obj.AMask: {Flags: gc.SizeW | gc.LeftRead | gc.RightWrite | gc.Move | gc.Conv | gc.PostInc},
	ppc64.AMOVHZ & obj.AMask: {Flags: gc.SizeW | gc.LeftRead | gc.RightWrite | gc.Move | gc.Conv},
	ppc64.AMOVW & obj.AMask:  {Flags: gc.SizeL | gc.LeftRead | gc.RightWrite | gc.Move | gc.Conv},

	ppc64.AISEL & obj.AMask: {Flags: gc.SizeQ | gc.RegRead | gc.From3Read | gc.RightWrite},

	// there is no AMOVWU.
	ppc64.AMOVWZU & obj.AMask: {Flags: gc.SizeL | gc.LeftRead | gc.RightWrite | gc.Move | gc.Conv | gc.PostInc},
	ppc64.AMOVWZ & obj.AMask:  {Flags: gc.SizeL | gc.LeftRead | gc.RightWrite | gc.Move | gc.Conv},
	ppc64.AMOVD & obj.AMask:   {Flags: gc.SizeQ | gc.LeftRead | gc.RightWrite | gc.Move},
	ppc64.AMOVDU & obj.AMask:  {Flags: gc.SizeQ | gc.LeftRead | gc.RightWrite | gc.Move | gc.PostInc},
	ppc64.AFMOVS & obj.AMask:  {Flags: gc.SizeF | gc.LeftRead | gc.RightWrite | gc.Move | gc.Conv},
	ppc64.AFMOVSX & obj.AMask: {Flags: gc.SizeF | gc.LeftRead | gc.RightWrite | gc.Move | gc.Conv},
	ppc64.AFMOVSZ & obj.AMask: {Flags: gc.SizeF | gc.LeftRead | gc.RightWrite | gc.Move | gc.Conv},
	ppc64.AFMOVD & obj.AMask:  {Flags: gc.SizeD | gc.LeftRead | gc.RightWrite | gc.Move},
	ppc64.AMFVSRD & obj.AMask: {Flags: gc.SizeD | gc.LeftRead | gc.RightWrite | gc.Move},
	ppc64.AMTVSRD & obj.AMask: {Flags: gc.SizeD | gc.LeftRead | gc.RightWrite | gc.Move},

	// Jumps
	ppc64.ABR & obj.AMask:  {Flags: gc.Jump | gc.Break},
	ppc64.ABL & obj.AMask:  {Flags: gc.Call},
	ppc64.ABVS & obj.AMask: {Flags: gc.Cjmp},
	ppc64.ABVC & obj.AMask: {Flags: gc.Cjmp},
	ppc64.ABEQ & obj.AMask: {Flags: gc.Cjmp},
	ppc64.ABNE & obj.AMask: {Flags: gc.Cjmp},
	ppc64.ABGE & obj.AMask: {Flags: gc.Cjmp},
	ppc64.ABLT & obj.AMask: {Flags: gc.Cjmp},
	ppc64.ABGT & obj.AMask: {Flags: gc.Cjmp},
	ppc64.ABLE & obj.AMask: {Flags: gc.Cjmp},
	obj.ARET:               {Flags: gc.Break},
	obj.ADUFFZERO:          {Flags: gc.Call},
	obj.ADUFFCOPY:          {Flags: gc.Call},
}

func initproginfo() {
	var addvariant = []int{V_CC, V_V, V_CC | V_V}

	// Perform one-time expansion of instructions in progtable to
	// their CC, V, and VCC variants
	for i := range progtable {
		as := obj.As(i)
		if progtable[as].Flags == 0 {
			continue
		}
		variant := as2variant(as)
		for i := range addvariant {
			as2 := variant2as(as, variant|addvariant[i])
			if as2 != 0 && progtable[as2&obj.AMask].Flags == 0 {
				progtable[as2&obj.AMask] = progtable[as]
			}
		}
	}
}

func proginfo(p *obj.Prog) gc.ProgInfo {
	info := progtable[p.As&obj.AMask]
	if info.Flags == 0 {
		gc.Fatalf("proginfo: unknown instruction %v", p)
	}

	if (info.Flags&gc.RegRead != 0) && p.Reg == 0 {
		info.Flags &^= gc.RegRead
		info.Flags |= gc.RightRead /*CanRegRead |*/
	}

	if p.From.Type == obj.TYPE_ADDR && p.From.Sym != nil && (info.Flags&gc.LeftRead != 0) {
		info.Flags &^= gc.LeftRead
		info.Flags |= gc.LeftAddr
	}

	return info
}

// Instruction variants table, populated by initvariants via Main.
// The index is the base form of the instruction, masked by obj.AMask.
// The 4 values are the unmasked base form, then the unmasked CC, V,
// and VCC variants, respectively.
var varianttable = [ppc64.ALAST & obj.AMask][4]obj.As{}

func initvariant(as obj.As, variants ...obj.As) {
	vv := &varianttable[as&obj.AMask]
	vv[0] = as
	for i, v := range variants {
		vv[i+1] = v
	}
}

func initvariants() {
	initvariant(ppc64.AADD, ppc64.AADDCC, ppc64.AADDV, ppc64.AADDVCC)
	initvariant(ppc64.AADDC, ppc64.AADDCCC, ppc64.AADDCV, ppc64.AADDCVCC)
	initvariant(ppc64.AADDE, ppc64.AADDECC, ppc64.AADDEV, ppc64.AADDEVCC)
	initvariant(ppc64.AADDME, ppc64.AADDMECC, ppc64.AADDMEV, ppc64.AADDMEVCC)
	initvariant(ppc64.AADDZE, ppc64.AADDZECC, ppc64.AADDZEV, ppc64.AADDZEVCC)
	initvariant(ppc64.AAND, ppc64.AANDCC)
	initvariant(ppc64.AANDN, ppc64.AANDNCC)
	initvariant(ppc64.ACNTLZD, ppc64.ACNTLZDCC)
	initvariant(ppc64.ACNTLZW, ppc64.ACNTLZWCC)
	initvariant(ppc64.ADIVD, ppc64.ADIVDCC, ppc64.ADIVDV, ppc64.ADIVDVCC)
	initvariant(ppc64.ADIVDU, ppc64.ADIVDUCC, ppc64.ADIVDUV, ppc64.ADIVDUVCC)
	initvariant(ppc64.ADIVW, ppc64.ADIVWCC, ppc64.ADIVWV, ppc64.ADIVWVCC)
	initvariant(ppc64.ADIVWU, ppc64.ADIVWUCC, ppc64.ADIVWUV, ppc64.ADIVWUVCC)
	initvariant(ppc64.AEQV, ppc64.AEQVCC)
	initvariant(ppc64.AEXTSB, ppc64.AEXTSBCC)
	initvariant(ppc64.AEXTSH, ppc64.AEXTSHCC)
	initvariant(ppc64.AEXTSW, ppc64.AEXTSWCC)
	initvariant(ppc64.AFABS, ppc64.AFABSCC)
	initvariant(ppc64.AFADD, ppc64.AFADDCC)
	initvariant(ppc64.AFADDS, ppc64.AFADDSCC)
	initvariant(ppc64.AFCFID, ppc64.AFCFIDCC)
	initvariant(ppc64.AFCFIDU, ppc64.AFCFIDUCC)
	initvariant(ppc64.AFCTID, ppc64.AFCTIDCC)
	initvariant(ppc64.AFCTIDZ, ppc64.AFCTIDZCC)
	initvariant(ppc64.AFCTIW, ppc64.AFCTIWCC)
	initvariant(ppc64.AFCTIWZ, ppc64.AFCTIWZCC)
	initvariant(ppc64.AFDIV, ppc64.AFDIVCC)
	initvariant(ppc64.AFDIVS, ppc64.AFDIVSCC)
	initvariant(ppc64.AFMADD, ppc64.AFMADDCC)
	initvariant(ppc64.AFMADDS, ppc64.AFMADDSCC)
	initvariant(ppc64.AFMOVD, ppc64.AFMOVDCC)
	initvariant(ppc64.AFMSUB, ppc64.AFMSUBCC)
	initvariant(ppc64.AFMSUBS, ppc64.AFMSUBSCC)
	initvariant(ppc64.AFMUL, ppc64.AFMULCC)
	initvariant(ppc64.AFMULS, ppc64.AFMULSCC)
	initvariant(ppc64.AFNABS, ppc64.AFNABSCC)
	initvariant(ppc64.AFNEG, ppc64.AFNEGCC)
	initvariant(ppc64.AFNMADD, ppc64.AFNMADDCC)
	initvariant(ppc64.AFNMADDS, ppc64.AFNMADDSCC)
	initvariant(ppc64.AFNMSUB, ppc64.AFNMSUBCC)
	initvariant(ppc64.AFNMSUBS, ppc64.AFNMSUBSCC)
	initvariant(ppc64.AFRES, ppc64.AFRESCC)
	initvariant(ppc64.AFRSP, ppc64.AFRSPCC)
	initvariant(ppc64.AFRSQRTE, ppc64.AFRSQRTECC)
	initvariant(ppc64.AFSEL, ppc64.AFSELCC)
	initvariant(ppc64.AFSQRT, ppc64.AFSQRTCC)
	initvariant(ppc64.AFSQRTS, ppc64.AFSQRTSCC)
	initvariant(ppc64.AFSUB, ppc64.AFSUBCC)
	initvariant(ppc64.AFSUBS, ppc64.AFSUBSCC)
	initvariant(ppc64.AMTFSB0, ppc64.AMTFSB0CC)
	initvariant(ppc64.AMTFSB1, ppc64.AMTFSB1CC)
	initvariant(ppc64.AMULHD, ppc64.AMULHDCC)
	initvariant(ppc64.AMULHDU, ppc64.AMULHDUCC)
	initvariant(ppc64.AMULHW, ppc64.AMULHWCC)
	initvariant(ppc64.AMULHWU, ppc64.AMULHWUCC)
	initvariant(ppc64.AMULLD, ppc64.AMULLDCC, ppc64.AMULLDV, ppc64.AMULLDVCC)
	initvariant(ppc64.AMULLW, ppc64.AMULLWCC, ppc64.AMULLWV, ppc64.AMULLWVCC)
	initvariant(ppc64.ANAND, ppc64.ANANDCC)
	initvariant(ppc64.ANEG, ppc64.ANEGCC, ppc64.ANEGV, ppc64.ANEGVCC)
	initvariant(ppc64.ANOR, ppc64.ANORCC)
	initvariant(ppc64.AOR, ppc64.AORCC)
	initvariant(ppc64.AORN, ppc64.AORNCC)
	initvariant(ppc64.AREM, ppc64.AREMCC, ppc64.AREMV, ppc64.AREMVCC)
	initvariant(ppc64.AREMD, ppc64.AREMDCC, ppc64.AREMDV, ppc64.AREMDVCC)
	initvariant(ppc64.AREMDU, ppc64.AREMDUCC, ppc64.AREMDUV, ppc64.AREMDUVCC)
	initvariant(ppc64.AREMU, ppc64.AREMUCC, ppc64.AREMUV, ppc64.AREMUVCC)
	initvariant(ppc64.ARLDC, ppc64.ARLDCCC)
	initvariant(ppc64.ARLDCL, ppc64.ARLDCLCC)
	initvariant(ppc64.ARLDCR, ppc64.ARLDCRCC)
	initvariant(ppc64.ARLDMI, ppc64.ARLDMICC)
	initvariant(ppc64.ARLWMI, ppc64.ARLWMICC)
	initvariant(ppc64.ARLWNM, ppc64.ARLWNMCC)
	initvariant(ppc64.ASLD, ppc64.ASLDCC)
	initvariant(ppc64.ASLW, ppc64.ASLWCC)
	initvariant(ppc64.ASRAD, ppc64.ASRADCC)
	initvariant(ppc64.ASRAW, ppc64.ASRAWCC)
	initvariant(ppc64.ASRD, ppc64.ASRDCC)
	initvariant(ppc64.ASRW, ppc64.ASRWCC)
	initvariant(ppc64.ASUB, ppc64.ASUBCC, ppc64.ASUBV, ppc64.ASUBVCC)
	initvariant(ppc64.ASUBC, ppc64.ASUBCCC, ppc64.ASUBCV, ppc64.ASUBCVCC)
	initvariant(ppc64.ASUBE, ppc64.ASUBECC, ppc64.ASUBEV, ppc64.ASUBEVCC)
	initvariant(ppc64.ASUBME, ppc64.ASUBMECC, ppc64.ASUBMEV, ppc64.ASUBMEVCC)
	initvariant(ppc64.ASUBZE, ppc64.ASUBZECC, ppc64.ASUBZEV, ppc64.ASUBZEVCC)
	initvariant(ppc64.AXOR, ppc64.AXORCC)

	for i := range varianttable {
		vv := &varianttable[i]
		if vv[0] == 0 {
			// Instruction has no variants
			varianttable[i][0] = obj.As(i)
			continue
		}

		// Copy base form to other variants
		if vv[0]&obj.AMask == obj.As(i) {
			for _, v := range vv {
				if v != 0 {
					varianttable[v&obj.AMask] = varianttable[i]
				}
			}
		}
	}
}

// as2variant returns the variant (V_*) flags of instruction as.
func as2variant(as obj.As) int {
	for i, v := range varianttable[as&obj.AMask] {
		if v&obj.AMask == as&obj.AMask {
			return i
		}
	}
	gc.Fatalf("as2variant: instruction %v is not a variant of itself", as&obj.AMask)
	return 0
}

// variant2as returns the instruction as with the given variant (V_*) flags.
// If no such variant exists, this returns 0.
func variant2as(as obj.As, flags int) obj.As {
	return varianttable[as&obj.AMask][flags]
}
