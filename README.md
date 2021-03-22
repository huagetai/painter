## 介绍
paiter是一个生成微信分享图的工具。
开发有翼小助手微信小程序时，需要生成转发分享图和朋友圈海报图的生成。
起初使用的是https://github.com/Kujiale-Mobile/Painter前端一个很强大的工具，
但是由于微信小程序前端缓存限制以及分享触发机制的规则，导致生成的分享图总是出问题，于是乎有了本项目。

## 功能
- 支持图片、文字、矩形等view
- 图片支持本地和网络图片，网络图片会启用本地缓存
- 支持的图片缩放裁剪模式丰富，有smart、scaleToFill、aspectFit、aspectFill、top、bottom、center、left、right、top left、top right、bottom left、bottom right。除smart外，其余和微信小程序组件image的一致
- 支持图片border-radis属性，圆角图片不再是问题
- 文字支持LineHeight、LineClamp属性，超出width自动截取
- 文字粗细程度，支持heavy、bold、bolder、normal、lighter

## 实体对象说明
### 调色板对象
```go
type Palette struct {
	Background string    // 调色板背景色,比如：#fffffff
	Width int            // 调色板的宽,比如：700
	Height int           // 调色板的高,比如：600
	BorderRadius string  // 边框圆角半径,比如：8
	Views []view.View    // VIEW切片
}
```
### VIEW对象

#### VIEW接口
```go
type View interface {
	Paint(ctx *PosterContext)
}
```
#### 图片VIEW
```go
type ImageView struct {
	X int                 // 横坐标,比如：0
	Y int                 // 纵坐标,比如：60
	Width int             // 图片宽,比如：60
	Height int            // 图片高,比如：60
	Mode string           // 缩放裁剪模式,支持：smart、scaleToFill、aspectFit、aspectFill、top、bottom、center、left、right、top left、top right、bottom left、bottom right
	URI string            // 图片URI，支持本地文件路径或网络地址。支持jpeg、png、gif、webp
	BorderRadis float64   // 边框圆角半径,比如：8,需要圆形头像时，取width的一半
}
```

#### 矩形VIEW
```go
type RectangleView struct {
	X float64              // 横坐标,比如：0
	Y float64              // 纵坐标,比如：60
	Width float64          // 矩形宽,比如：60
	Height float64         // 矩形高,比如：60
	BorderRadis float64    // 边框圆角半径,比如：8,需要圆形头像时，取width的一半
	BackgroudColor string  // 填充颜色，比如：#888888
}
```

#### 文字VIEW
```go
type TextView struct {
	X float64              // 横坐标,比如：0
	Y float64              // 纵坐标,比如：60
	Width float64          // 文字最大宽,比如：60
	LineHeight float64     // 文字行高，小于字体高度时，自动为字体高度的1.25倍
	LineClamp string       // 行数
	Text string            // 文字，目前使用阿里普惠体，不支持表情符号
	FontSize float64       // 文字大小
	FontWeight string      // 文字粗细程度，支持heavy、bold、bolder、normal、lighter
	Color string           // 文字颜色，比如：#888888
	Align string           // 文字对齐方式，支持left、center、right
}
```

## 使用说明

```bash
go get github.com/huagetai/painter
```


```go
    // 构造views
    var views []view.View
    //...
    // 构造调色板
    p := palette.Palette {
		Background: "#FFFFFF",
		Width: 700,
		Height: 600,
		Views: views,
    }
    // 构造io.Writer 
    f,_ := os.Create("./assets/test.png")
    w := bufio.NewWriter(f)
    p.Paint(&w)
```

painter.go是一个生成微信分享图片的样例


## 使用的开源项目
- https://github.com/disintegration/imaging
- https://github.com/muesli/smartcrop
- https://github.com/nfnt/resize
- https://github.com/fogleman/gg
