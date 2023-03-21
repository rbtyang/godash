/*
@Editor robotyang at 2023

AES单测运行说明：
    1、先在 godash 项目根目录下，执行 yarn 命令，安装前端依赖包
    2、运行单测的方式：
        方法1、在当前文件 右键Run（如果是 goland 想执行 该JS测试用例，则需要安装 Node.js IDE插件）
        方法2、点击IDE的单测运行图标（正常是在 在describe左边行数栏的 一个Run图标）
        方法3、在 godash 项目根目录，执行 `yarn run test`
*/
import aes from './aes.js'
import assert from 'assert'

describe('aes', function () {
    describe('encrypt', function () {
        it('mixed plaintext 1', function () {
            const ciphertext = aes.Encrypt('123123阿斯蒂芬!@#asdasd', 'ABCDEF1234123412')
            console.log('       > ciphertext=', ciphertext);
            assert.equal(ciphertext, "9e2819811c90f03c407ecfc7253b240556ff169e3371127c46a651ae1920d8df")
        })
        it('mixed plaintext 2 diff secret', function () {
            const ciphertext = aes.Encrypt('123123阿斯蒂芬!@#asdasd', 'M8xMxeX6rgBsveTF')
            console.log('       > ciphertext=', ciphertext);
            assert.equal(ciphertext, "3c258944eb76163936f7b01aa453425f28110783c85d6241b319aa24c25e9023")
        })
        it('mixed plaintext 3 long secret', function () {
            const ciphertext = aes.Encrypt('123123阿斯蒂芬!@#asdasd', 'M8xMxeX6rgBsveTFM8xMxeX6rgBsveTF')
            console.log('       > ciphertext=', ciphertext);
            assert.equal(ciphertext, "3c258944eb76163936f7b01aa453425f28110783c85d6241b319aa24c25e9023")
        })
    })
})
