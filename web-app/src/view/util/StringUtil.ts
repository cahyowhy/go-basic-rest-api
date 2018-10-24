/**
 * Created by fajar on 06/28/2018.
 * Updated by cahyo on 16/08/18.
 */

/**
 * isEmail check match string is email
 *
 * @param {string} param string to be check
 * @returns {boolean}
 */
export const isEmail = (param: string = '') => {
    const regex: any = /^(([^<>()[\]\\.,;:\s@\"]+(\.[^<>()[\]\\.,;:\s@\"]+)*)|(\".+\"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;

    return regex.test(param);
};

/**
 * convert camelcase text into valid format text / readable
 * use it to convert table column name into valid format text
 *
 * @param {string} param
 * @returns {string}
 */
export const normalizeCC = (param: string): string => {
    const result = param.replace(/([A-Z])/g, " $1");
    return result.charAt(0).toUpperCase() + result.slice(1);
};

/**
 * convert camelcase text into valid format text / readable
 * use it to convert table column name into valid format text
 *
 * @param {string} param
 * @returns {string}
 */
export const normalizeUnderscore = (param: string): string => {
    const result = param.replace(/_/g, " ");
    return result.charAt(0).toUpperCase() + result.slice(1);
};

/**
 * convert text with space into camelcase
 *
 * @param {string} str
 * @returns {string}
 */
export const camelize = (str: string): string => {
    return str.replace(/(?:^\w|[A-Z]|\b\w|\s+)/g, function (match, index) {
        if (+match === 0) return "";
        return index == 0 ? match.toLowerCase() : match.toUpperCase();
    });
};

/**
 * generate guid
 *
 * @returns {string}
 */
export const generateGUID = () => {
    const s4 = () => {
        return Math.floor((1 + Math.random()) * 0x10000)
            .toString(16)
            .substring(1);
    };

    return s4() + s4() + '-' + s4() + '-' + s4() + '-' +
        s4() + '-' + s4() + s4() + s4();
};

export const htmlTagToText = (htmlTag: string) => {
    const encode = (param: string) => {
        return param.replace(/&#(\d+);/g, function (match: any, dec: any) {
            return String.fromCharCode(dec);
        });
    };

    return encode(htmlTag.replace(/<[^>]+>/g, ''));
};
