Doc
==============
https://doc.praxiscashier.com/integration_docs/latest/cashier_api/cashier


Env
==============
sandbox: https://pci-gw-test.praxispay.com/cashier/cashier
Live: https://gw.praxisgate.com/cashier/cashier


鉴权
==============
对请求参数签名后,放到了Gt-Authentication 这个header里


回调地址
==============
在下单时，参数里传递进来的(动态可修改)


Comment
===============
1. both support deposit && withdrawl
2. deposit和withdrawl是同一个接口，只是通过参数来区分是哪种行为.
3. 在pre接口中有一个参数notification_url, 这个是充值回调url指定
4. pre接口是需要签名的, 是几个字段的hash值计算(包含私有的app_key字段) https://doc.praxiscashier.com/integration_docs/latest/overview/authentication 