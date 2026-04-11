from zipfile import ZipFile

from PIL import Image, ImageDraw, ImageFont


def main():
    with ZipFile("04b_03.zip") as zf:
        with zf.open("04B_03__.TTF") as fp:
            font = ImageFont.truetype(fp, size=8)

    img = Image.new("1", (16*8, 6*8), 255)
    d = ImageDraw.Draw(img)

    for o in range(33, 127):
        x = (o & 15) << 3
        y = ((o & 0xf0) >> 1) - 16
        d.text((x, y), chr(o), fill=0, font=font)

    img.save("04b_03.png")


if __name__ == "__main__":
    main()
