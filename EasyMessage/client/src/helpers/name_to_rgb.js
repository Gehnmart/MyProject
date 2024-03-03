import tinycolor from "tinycolor2";

export default function name_to_rgb(name) {
    const shift = 60;
    const rgb = {
        r: (name.charCodeAt(0) + shift) % 255 + 15,
        g: (name.charCodeAt(1) + shift) % 255 + 15,
        b: (name.length * 15) % 255 - 15,
    }
    return [tinycolor(rgb).darken(10).toHexString(), tinycolor(rgb).lighten(20).toHexString()];
}