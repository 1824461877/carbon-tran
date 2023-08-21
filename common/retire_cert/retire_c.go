package retire_cert

import (
	"fmt"
	"github.com/fogleman/gg"
	"image/color"
	"path"
)

func (c *RetireC) build() RetireConfigInter {
	bgImage, err := gg.LoadImage(c.retireConfig.ImageBackground)
	if err != nil {
		return nil
	}
	imgWidth := bgImage.Bounds().Dx()
	imgHeight := bgImage.Bounds().Dy()
	dc := gg.NewContext(imgWidth, imgHeight)
	dc.DrawImage(bgImage, 0, 0)
	c.x = float64(imgWidth / 2)
	c.dc = dc
	return c
}

func (c *RetireC) name(projectName string, onTime string) string {
	return RCID(projectName, onTime)
}

func (c *RetireC) savePath(name string) string {
	return path.Join(c.retireConfig.Path, name)
}

func (c *RetireC) Create(rc *RetireCertificate) (fileName string, savePath string, errs error) {
	var (
		err error
	)
	c.dc.SetColor(color.RGBA{R: 90, G: 90, B: 90, A: 255})
	if err = c.dc.LoadFontFace(c.retireConfig.FontTTF, 24); err != nil {
		return "", "", err
	}

	c.dc.DrawStringAnchored("我们很高兴确认", c.x, 395, 0.5, 0.5)
	c.dc.DrawStringAnchored("该 credit 是代表个人或私人注销。", c.x, 660, 0.5, 0.5)
	c.dc.DrawStringAnchored(fmt.Sprintf("Retired by %v", rc.RetiredBy), c.x, 715, 0.5, 0.5)
	c.dc.DrawStringAnchored(fmt.Sprintf("Project: %v", rc.ProjectName), c.x, 770, 0.5, 0.5)
	c.dc.DrawStringAnchored(fmt.Sprintf("这样 credit 已经注销，节省了 %v 吨二氧化碳排放量", rc.VerifiedNumber), c.x, 840, 0.5, 0.5)
	c.dc.DrawStringAnchored("防止释放到大气中。", c.x, 880, 0.5, 0.5)
	c.dc.DrawStringAnchored("感谢你们为一个更安全的气候和更可持续的世界投资。", c.x, 920, 0.5, 0.5)

	if err = c.dc.LoadFontFace(c.retireConfig.FontTTF, 40); err != nil {
		return "", "", err
	}
	c.dc.SetColor(color.RGBA{R: 0, G: 187, B: 191, A: 255})
	c.dc.DrawStringAnchored(rc.ByFrom, c.x, 460, 0.5, 0.5)
	c.dc.DrawStringAnchored(rc.ByTo, c.x, 545, 0.5, 0.5)

	if err = c.dc.LoadFontFace(c.retireConfig.FontTTF, 20); err != nil {
		return "", "", err
	}
	c.dc.SetColor(color.RGBA{R: 90, G: 90, B: 90, A: 255})
	c.dc.DrawStringAnchored("by", c.x, 500, 0.5, 0.5)
	c.dc.DrawStringAnchored(fmt.Sprintf("on %v", rc.OnTime), c.x, 595, 0.5, 0.5)

	name := c.name(rc.ProjectName, rc.OnTime)
	err = c.dc.SavePNG(path.Join(c.retireConfig.Path, name+".png"))
	if err != nil {
		return "", "", err
	}

	return name, c.savePath(name), nil
}

func NewRetireC(uploadPath string, imageBackground string, fontTTF string) RetireConfigInter {
	rc := &RetireC{
		retireConfig: RetireConfig{
			Path:            uploadPath,
			ImageBackground: imageBackground,
			FontTTF:         fontTTF,
		},
	}

	return rc.build()
}
