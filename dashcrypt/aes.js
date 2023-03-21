/*
@Editor robotyang at 2023

AES包引入使用说明：
    1、先在前端项目根目录下，执行 yarn add crypto-js 安装 https://github.com/brix/crypto-js
    2、再将当前的 js文件 copy到 前端项目中的任意位置，如 @/utils/aes.js
    3、然后在 需要的地方 进行引入 并使用，如下：
        import aes from '@/utils/aes.js'
        const ciphertext = aes.Encrypt('123123阿斯蒂芬!@#asdasd', 'ABCDEF1234123412')
*/
import CryptoJS from 'crypto-js'

// @param string plaintext 需要被对称加密的明文
// @param string secret 对称加密的密钥  必须是16长度,为了和后端交互 key字符串必须是16进制字符串,否在给golang进行string -> []byte带来困难
// @reference https://mojotv.cn/go/crypto-js-with-golang
const Encrypt = (plaintext, secret) => {
    secret = _paddingLeft(secret.slice(0, 16), 16)// 保证key的长度为16byte,进行'0'补位
    secret = CryptoJS.enc.Utf8.parse(secret)
    // 加密结果返回的是CipherParams object类型
    // key 和 iv 使用同一个值
    var encrypted = CryptoJS.AES.encrypt(plaintext, secret, {
        iv: secret,
        mode: CryptoJS.mode.CBC, // CBC算法
        padding: CryptoJS.pad.Pkcs7 // 使用pkcs7 进行padding 后端需要注意
    })
    // ciphertext是密文,toString()内传编码格式,比如Base64,这里用了16进制
    // 如果密文要放在 url的参数中 建议进行 base64-url-encoding 和 hex encoding, 不建议使用base64 encoding
    return encrypted.ciphertext.toString(CryptoJS.enc.Hex) // 后端必须进行相反操作
}

// @param string secret 对称加密的密钥  确保key的长度,使用 0 字符来补位
// @param integer length 建议 16 24 32
const _paddingLeft = (secret, length) => {
    let pkey = secret.toString()
    const l = pkey.length
    if (l < length) {
        pkey = new Array(length - l + 1).join('0') + pkey
    } else if (l > length) {
        pkey = pkey.slice(length)
    }
    return pkey
}

export default {
    Encrypt
}