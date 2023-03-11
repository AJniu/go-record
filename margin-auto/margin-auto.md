#### margin: 0 auto; 为何能水平居中？
margin: 0 auto; 中的 auto 是指 `平分剩余空间`。
比如宽度为200,父元素的宽度为1000,那么auto就是指水平方向平分剩余的宽度(1000-200)/2)


##### 单独设置margin-left 或 margin-right效果
如果单独设置 margin-left: auto; 则会实现与 float: left;相同效果。因为默认margin-right为0，所以除子元素自身宽度外，所有剩余空间都会被认为是margin-left的剩余空间。

### 使用margin: auto;实现水平垂直居中。
1. 使用flex布局：父元素设置为displa: flex; 子元素: margin: auto; 即可实现水平垂直居中。
2. 使用position绝对定位： 将子元素定位设置为absolute，且上下左右全部设置为0，再margin: auto就可实现。 -- 原理就是当上下左右全部为0时，子元素就会填充整个父元素的所有可用空间。那么auto就有了作用，会`平均分配剩余空间`。
