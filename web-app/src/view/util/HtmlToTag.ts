/**
 * Created by cahyo on 09/04/2017.
 */

/**
 * Converts a string to its html characters completely.
 * @param param
 **/
export function decode(param): string {
    let buf = [];
    for (let i = param.length - 1; i >= 0; i--) {
        buf.unshift(['&#', param[i].charCodeAt(), ';'].join(''));
    }

    return buf.join('');
}

/**
 * Converts an html characterSet into its original character.
 * @param {String} param htmlSet entities
 **/
export function encode(param): string {
    return param.replace(/&#(\d+);/g, function (match, dec) {
        return String.fromCharCode(dec);
    });
}

/**
 * pick only string from html string
 * @param param
 */
export default function (param): string {
    return encode(param.replace(/<[^>]+>/g, ''));
}
