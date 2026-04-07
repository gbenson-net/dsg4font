from textwrap import wrap

from PIL import Image, ImageChops


def main():
    src = Image.open("04b_08.png").convert("L")
    src = ImageChops.invert(src)
    glyphs = [0]
    for o in range(33, 128):
        x = (o & 15) << 3
        y = ((o & 0xf0) >> 1) - 16
        img = src.crop((x, y, x+8, y+8))
        l, t, r, b = img.getbbox()
        w = r - l
        h = b - t
        assert l == 0 or o == 33
        assert t >= 1
        assert 0 < w <= 5
        assert 0 < h <= 5
        img = img.crop((l, 1, l+5, 6))
        assert img.size == (5, 5)
        v = 0
        for x in range(4, -1, -1):
            for y in range(5):
                v <<= 1
                if img.getpixel((x, y)):
                    v |= 1
        glyphs.append(v)
    lines = wrap(" ".join(f"{v}," for v in glyphs), width=79)
    print("""\
package dsg4font

// Face04B08 is a Face based on the 04B-08 TrueType font by Yuji Oshimoto.
// (http://www.dsg4.com/04/).
//
// At the moment, it holds the printable characters in ASCII starting with
// space, and the Unicode replacement character U+FFFD.  Most glyphs are 6
// pixels square, though some are narrower.
//
// Its data is entirely self-contained and does not require loading from
// separate files.
var Face04B08 = &Face55vw{""")
    print("\n".join(f"\t{line}" for line in lines))
    print("}")


if __name__ == "__main__":
    main()
