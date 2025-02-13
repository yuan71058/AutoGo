package opencv

func (c ColorConversionCode) String() string {
	switch c {
	case ColorBGRToBGRA:
		return "color-bgr-to-bgra"
	case ColorBGRAToBGR:
		return "color-bgra-to-bgr"
	case ColorBGRToRGBA:
		return "color-bgr-to-rgba"
	case ColorRGBAToBGR:
		return "color-rgba-to-bgr"
	case ColorBGRToRGB:
		return "color-bgr-to-rgb"
	case ColorBGRAToRGBA:
		return "color-bgra-to-rgba"
	case ColorBGRToGray:
		return "color-bgr-to-gray"
	case ColorRGBToGray:
		return "color-rgb-to-gray"
	case ColorGrayToBGR:
		return "color-gray-to-bgr"
	case ColorGrayToBGRA:
		return "color-gray-to-bgra"
	case ColorBGRAToGray:
		return "color-bgra-to-gray"
	case ColorRGBAToGray:
		return "color-rgba-to-gray"
	case ColorBGRToBGR565:
		return "color-bgr-to-bgr565"
	case ColorRGBToBGR565:
		return "color-rgb-to-bgr565"
	case ColorBGR565ToBGR:
		return "color-bgr565-to-bgr"
	case ColorBGR565ToRGB:
		return "color-bgr565-to-rgb"
	case ColorBGRAToBGR565:
		return "color-bgra-to-bgr565"
	case ColorRGBAToBGR565:
		return "color-rgba-to-bgr565"
	case ColorBGR565ToBGRA:
		return "color-bgr565-to-bgra"
	case ColorBGR565ToRGBA:
		return "color-bgr565-to-rgba"
	case ColorGrayToBGR565:
		return "color-gray-to-bgr565"
	case ColorBGR565ToGray:
		return "color-bgr565-to-gray"
	case ColorBGRToBGR555:
		return "color-bgr-to-bgr555"
	case ColorRGBToBGR555:
		return "color-rgb-to-bgr555"
	case ColorBGR555ToBGR:
		return "color-bgr555-to-bgr"
	case ColorBGRAToBGR555:
		return "color-bgra-to-bgr555"
	case ColorRGBAToBGR555:
		return "color-rgba-to-bgr555"
	case ColorBGR555ToBGRA:
		return "color-bgr555-to-bgra"
	case ColorBGR555ToRGBA:
		return "color-bgr555-to-rgba"
	case ColorGrayToBGR555:
		return "color-gray-to-bgr555"
	case ColorBGR555ToGRAY:
		return "color-bgr555-to-gray"
	case ColorBGRToXYZ:
		return "color-bgr-to-xyz"
	case ColorRGBToXYZ:
		return "color-rgb-to-xyz"
	case ColorXYZToBGR:
		return "color-xyz-to-bgr"
	case ColorXYZToRGB:
		return "color-xyz-to-rgb"
	case ColorBGRToYCrCb:
		return "color-bgr-to-ycrcb"
	case ColorRGBToYCrCb:
		return "color-rgb-to-ycrcb"
	case ColorYCrCbToBGR:
		return "color-ycrcb-to-bgr"
	case ColorYCrCbToRGB:
		return "color-ycrcb-to-rgb"
	case ColorBGRToHSV:
		return "color-bgr-to-hsv"
	case ColorRGBToHSV:
		return "color-rgb-to-hsv"
	case ColorBGRToLab:
		return "color-bgr-to-lab"
	case ColorRGBToLab:
		return "color-rgb-to-lab"
	case ColorBGRToLuv:
		return "color-bgr-to-luv"
	case ColorRGBToLuv:
		return "color-rgb-to-luv"
	case ColorBGRToHLS:
		return "color-bgr-to-hls"
	case ColorRGBToHLS:
		return "color-rgb-to-hls"
	case ColorHSVToBGR:
		return "color-hsv-to-bgr"
	case ColorHSVToRGB:
		return "color-hsv-to-rgb"
	case ColorLabToBGR:
		return "color-lab-to-bgr"
	case ColorLabToRGB:
		return "color-lab-to-rgb"
	case ColorLuvToBGR:
		return "color-luv-to-bgr"
	case ColorLuvToRGB:
		return "color-luv-to-rgb"
	case ColorHLSToBGR:
		return "color-hls-to-bgr"
	case ColorHLSToRGB:
		return "color-hls-to-rgb"
	case ColorBGRToHSVFull:
		return "color-bgr-to-hsv-full"
	case ColorRGBToHSVFull:
		return "color-rgb-to-hsv-full"
	case ColorBGRToHLSFull:
		return "color-bgr-to-hls-full"
	case ColorRGBToHLSFull:
		return "color-rgb-to-hls-full"
	case ColorHSVToBGRFull:
		return "color-hsv-to-bgr-full"
	case ColorHSVToRGBFull:
		return "color-hsv-to-rgb-full"
	case ColorHLSToBGRFull:
		return "color-hls-to-bgr-full"
	case ColorHLSToRGBFull:
		return "color-hls-to-rgb-full"
	case ColorLBGRToLab:
		return "color-lbgr-to-lab"
	case ColorLRGBToLab:
		return "color-lrgb-to-lab"
	case ColorLBGRToLuv:
		return "color-lbgr-to-luv"
	case ColorLRGBToLuv:
		return "color-lrgb-to-luv"
	case ColorLabToLBGR:
		return "color-lab-to-lbgr"
	case ColorLabToLRGB:
		return "color-lab-to-lrgb"
	case ColorLuvToLBGR:
		return "color-luv-to-lbgr"
	case ColorLuvToLRGB:
		return "color-luv-to-lrgb"
	case ColorBGRToYUV:
		return "color-bgr-to-yuv"
	case ColorRGBToYUV:
		return "color-rgb-to-yuv"
	case ColorYUVToBGR:
		return "color-yuv-to-bgr"
	case ColorYUVToRGB:
		return "color-yuv-to-rgb"

	case ColorYUVToRGBNV12:
		return "color-yuv-to-rgbnv12"
	case ColorYUVToBGRNV12:
		return "color-yuv-to-bgrnv12"
	case ColorYUVToRGBNV21:
		return "color-yuv-to-rgbnv21"
	case ColorYUVToBGRNV21:
		return "color-yuv-to-bgrnv21"

	case ColorYUVToRGBANV12:
		return "color-yuv-to-rgbanv12"
	case ColorYUVToBGRANV12:
		return "color-yuv-to-bgranv12"
	case ColorYUVToRGBANV21:
		return "color-yuv-to-rgbanv21"
	case ColorYUVToBGRANV21:
		return "color-yuv-to-bgranv21"

	case ColorYUVToRGBYV12:
		return "color-yuv-to-rgbyv12"
	case ColorYUVToBGRYV12:
		return "color-yuv-to-bgryv12"

	case ColorYUVToRGBIYUV:
		return "color-yuv-to-rgbiyuv"
	case ColorYUVToBGRIYUV:
		return "color-yuv-to-bgriyuv"

	case ColorYUVToRGBAYV12:
		return "color-yuv-to-rgbayv12"
	case ColorYUVToBGRAYV12:
		return "color-yuv-to-bgrayv12"
	case ColorYUVToRGBAIYUV:
		return "color-yuv-to-rgbaiyuv"
	case ColorYUVToBGRAIYUV:
		return "color-yuv-to-bgraiyuv"

	case ColorYUVToGRAY420:
		return "color-yuv-to-gray420"

	case ColorYUVToRGBUYVY:
		return "color-yuv-to-rgbuyvy"
	case ColorYUVToBGRUYVY:
		return "color-yuv-to-bgruyvy"

	case ColorYUVToRGBAUYVY:
		return "color-yuv-to-rgbauyvy"
	case ColorYUVToBGRAUYVY:
		return "color-yuv-to-bgrauyvy"

	case ColorYUVToRGBYUY2:
		return "color-yuv-to-rgbyuy2"
	case ColorYUVToBGRYUY2:
		return "color-yuv-to-bgryuy2"

	case ColorYUVToRGBYVYU:
		return "color-yuv-to-rgbyvyu"
	case ColorYUVToBGRYVYU:
		return "color-yuv-to-bgryvyu"

	case ColorYUVToRGBAYUY2:
		return "color-yuv-to-rgbayuy2"
	case ColorYUVToBGRAYUY2:
		return "color-yuv-to-bgrayuy2"

	case ColorYUVToRGBAYVYU:
		return "color-yuv-to-rgbayvyu"
	case ColorYUVToBGRAYVYU:
		return "color-yuv-to-bgrayvyu"

	case ColorYUVToGRAYUYVY:
		return "color-yuv-to-grayuyvy"
	case ColorYUVToGRAYYUY2:
		return "color-yuv-to-grayyuy2"

	case ColorRGBATomRGBA:
		return "color-rgba-to-mrgba"
	case ColormRGBAToRGBA:
		return "color-mrgba-to-rgba"

	case ColorRGBToYUVI420:
		return "color-rgb-to-yuvi420"
	case ColorBGRToYUVI420:
		return "color-bgr-to-yuvi420"

	case ColorRGBAToYUVI420:
		return "color-rgba-to-yuvi420"

	case ColorBGRAToYUVI420:
		return "color-bgra-to-yuvi420"
	case ColorRGBToYUVYV12:
		return "color-rgb-to-yuvyv12"
	case ColorBGRToYUVYV12:
		return "color-bgr-to-yuvyv12"
	case ColorRGBAToYUVYV12:
		return "color-rgba-to-yuvyv12"
	case ColorBGRAToYUVYV12:
		return "color-bgra-to-yuvyv12"

	case ColorBayerBGToBGR:
		return "color-bayer-bgt-to-bgr"
	case ColorBayerGBToBGR:
		return "color-bayer-gbt-to-bgr"
	case ColorBayerRGToBGR:
		return "color-bayer-rgt-to-bgr"
	case ColorBayerGRToBGR:
		return "color-bayer-grt-to-bgr"

	case ColorBayerBGToGRAY:
		return "color-bayer-bgt-to-gray"
	case ColorBayerGBToGRAY:
		return "color-bayer-gbt-to-gray"
	case ColorBayerRGToGRAY:
		return "color-bayer-rgt-to-gray"
	case ColorBayerGRToGRAY:
		return "color-bayer-grt-to-gray"

	case ColorBayerBGToBGRVNG:
		return "color-bayer-bgt-to-bgrvng"
	case ColorBayerGBToBGRVNG:
		return "color-bayer-gbt-to-bgrvng"
	case ColorBayerRGToBGRVNG:
		return "color-bayer-rgt-to-bgrvng"
	case ColorBayerGRToBGRVNG:
		return "color-bayer-grt-to-bgrvng"

	case ColorBayerBGToBGREA:
		return "color-bayer-bgt-to-bgrea"
	case ColorBayerGBToBGREA:
		return "color-bayer-gbt-to-bgrea"
	case ColorBayerRGToBGREA:
		return "color-bayer-rgt-to-bgrea"
	case ColorBayerGRToBGREA:
		return "color-bayer-grt-to-bgrea"

	case ColorBayerBGToBGRA:
		return "color-bayer-bgt-to-bgra"
	case ColorBayerGBToBGRA:
		return "color-bayer-gbt-to-bgra"
	case ColorBayerRGToBGRA:
		return "color-bayer-rgt-to-bgra"
	case ColorBayerGRToBGRA:
		return "color-bayer-grt-to-bgra"
	case ColorCOLORCVTMAX:
		return "color-color-cvt-max"
	}
	return ""
}
