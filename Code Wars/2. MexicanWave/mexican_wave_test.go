package mexican_wave

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Example test cases:", func() {
	It("should return the correct values", func() {
		Expect(wave(" x yz")).To(Equal([]string{" X yz", " x Yz", " x yZ"}))
		Expect(wave("abc")).To(Equal([]string{"Abc", "aBc", "abC"}))
		Expect(wave("abc")).To(Equal([]string{"Abc", "aBc", "abC"}))
		Expect(wave(" ab  c")).To(Equal([]string{" Ab  c", " aB  c", " ab  C"}))
		Expect(wave("")).To(Equal([]string{}))
		Expect(wave("z")).To(Equal([]string{"Z"}))
		Expect(wave("a a a a a")).To(Equal([]string{"A a a a a", "a A a a a", "a a A a a", "a a a A a", "a a a a A"}))
		Expect(wave("aaaaa")).To(Equal([]string{"Aaaaa", "aAaaa", "aaAaa", "aaaAa", "aaaaA"}))
		Expect(wave("                                                           ")).To(Equal([]string{}))
		Expect(wave(" a  b     c  defghijk l  m no pqrs tuvwx yz     ")).To(Equal([]string{" A  b     c  defghijk l  m no pqrs tuvwx yz     ", " a  B     c  defghijk l  m no pqrs tuvwx yz     ", " a  b     C  defghijk l  m no pqrs tuvwx yz     ", " a  b     c  Defghijk l  m no pqrs tuvwx yz     ", " a  b     c  dEfghijk l  m no pqrs tuvwx yz     ", " a  b     c  deFghijk l  m no pqrs tuvwx yz     ", " a  b     c  defGhijk l  m no pqrs tuvwx yz     ", " a  b     c  defgHijk l  m no pqrs tuvwx yz     ", " a  b     c  defghIjk l  m no pqrs tuvwx yz     ", " a  b     c  defghiJk l  m no pqrs tuvwx yz     ", " a  b     c  defghijK l  m no pqrs tuvwx yz     ", " a  b     c  defghijk L  m no pqrs tuvwx yz     ", " a  b     c  defghijk l  M no pqrs tuvwx yz     ", " a  b     c  defghijk l  m No pqrs tuvwx yz     ", " a  b     c  defghijk l  m nO pqrs tuvwx yz     ", " a  b     c  defghijk l  m no Pqrs tuvwx yz     ", " a  b     c  defghijk l  m no pQrs tuvwx yz     ", " a  b     c  defghijk l  m no pqRs tuvwx yz     ", " a  b     c  defghijk l  m no pqrS tuvwx yz     ", " a  b     c  defghijk l  m no pqrs Tuvwx yz     ", " a  b     c  defghijk l  m no pqrs tUvwx yz     ", " a  b     c  defghijk l  m no pqrs tuVwx yz     ", " a  b     c  defghijk l  m no pqrs tuvWx yz     ", " a  b     c  defghijk l  m no pqrs tuvwX yz     ", " a  b     c  defghijk l  m no pqrs tuvwx Yz     ", " a  b     c  defghijk l  m no pqrs tuvwx yZ     "}))
	})
})
