package main

import (
	"encoding/base64"
	"fmt"
	"strings"
)

func main() {
	// Define the base64 string
	// base64Str := `JVBERi0xLjQKJeLjz9MKMSAwIG9iago8PC9Nb2REYXRlKEQ6MjAyNDA5MTgwNjUwNDcrMDInMDAnKS9DcmVhdGlvbkRhdGUoRDoyMDI0MDkxODA2NTA0NyswMicwMCcpL1Byb2R1Y2VyKGlUZXh0riA1LjUuNyCpMjAwMC0yMDE1IGlUZXh0IEdyb3VwIE5WIFwoQUdQTC12ZXJzaW9uXCk7IG1vZGlmaWVkIHVzaW5nIGlUZXh0riA1LjUuNyCpMjAwMC0yMDE1IGlUZXh0IEdyb3VwIE5WIFwoQUdQTC12ZXJzaW9uXCkpPj4KZW5kb2JqCjIgMCBvYmoKPDwvRmlsdGVyL0ZsYXRlRGVjb2RlL0xlbmd0aCA0OD4-c3RyZWFtCnicK-TiKlQw0LNUMABBIG1oomdgbGhqYKJgZKhnZGZgYGCmkJzLVcgVCMQAzzkJDwplbmRzdHJlYW0KZW5kb2JqCjMgMCBvYmoKPDwvRmlsdGVyL0ZsYXRlRGVjb2RlL0xlbmd0aCAxNj4-c3RyZWFtCnicUwjkKuTiCgRDAA4gAhIKZW5kc3RyZWFtCmVuZG9iago0IDAgb2JqCjw8L1N1YnR5cGUvVHlwZTEvVHlwZS9Gb250L0Jhc2VGb250L0hlbHZldGljYS1Cb2xkL0VuY29kaW5nL1dpbkFuc2lFbmNvZGluZz4-CmVuZG9iago1IDAgb2JqCjw8L1N1YnR5cGUvVHlwZTEvVHlwZS9Gb250L0Jhc2VGb250L1RpbWVzLVJvbWFuL0VuY29kaW5nL1dpbkFuc2lFbmNvZGluZz4-CmVuZG9iago2IDAgb2JqCjw8L1N1YnR5cGUvRm9ybS9GaWx0ZXIvRmxhdGVEZWNvZGUvVHlwZS9YT2JqZWN0L01hdHJpeFsxIDAgMCAxIDAgMF0vRm9ybVR5cGUgMS9SZXNvdXJjZXM8PC9Gb250PDwvRjMgNyAwIFI-Pj4-L0JCb3hbMCAwIDE3NS4zNSA2OS42Nl0vTGVuZ3RoIDI5MD4-c3RyZWFtCniclVIxUgQxDOvzipTQGCuJ46Rl5ihp2IIPcAUzNPy_IHC7uczaDa1GkSJLHDvVGhMhVo7fHyET5IaBWHawUjYYQGKZyFR39E9pRzt1OTulQTDM1KkY0TyszD_LNFpApWYkS_5VOmuWSsmCnbL9vQyrs5MoqXMnODmrHDlXZj9OeidqomLdtc6gy_vG1Bw0Tf9FoenstIzQN7Cn49ALs6tzVDAfV12L5uxkBTdvPwBTcmYBscUCbX7i_l2MtXRDTepVizEiGxhjMTYxcjvKXRUKZr0ruYiXWdjNPGaTHbKSIzyGo4Zaq9cmFK6d5nniRUNlfoL_iV3D8xaeXnJscbsGDJwjolRSHcaDuH2Fh9d3weUtNfTO_Lh9hssWfgD_mNx0CmVuZHN0cmVhbQplbmRvYmoKNyAwIG9iago8PC9TdWJ0eXBlL1R5cGUxL1R5cGUvRm9udC9CYXNlRm9udC9IZWx2ZXRpY2EvRW5jb2RpbmcvV2luQW5zaUVuY29kaW5nPj4KZW5kb2JqCjggMCBvYmoKPDwvU3VidHlwZS9Gb3JtL0ZpbHRlci9GbGF0ZURlY29kZS9UeXBlL1hPYmplY3QvTWF0cml4WzEgMCAwIDEgMCAwXS9Gb3JtVHlwZSAxL1Jlc291cmNlczw8L0ZvbnQ8PC9GMyA3IDAgUj4-Pj4vQkJveFswIDAgMTkyLjczIDc3LjE2XS9MZW5ndGggMjU5Pj5zdHJlYW0KeJx1kz1OBTEMhPucIiU0g-3EP2mR4AS5Aq9AouH-BSnQSo83aAtL4_Eo_pKVvhDRFeY9Et6_P5qd8qj-agO5Lk0Fwrx6Iqh-6iI5ZojJTpKIIv5hSCP-4TCmT4El0x1DSb4LRjAGDnZMT7CtwrAYtVPXIjE5ICw-HU71uvx3-eWIQfQlCEZhOae5itNUGbCgz6E4T1UH21i1wDZQc45I7QzQxpBr4j5qFLzYxFQIbwTUWJQrNGkjYMqiQi9SfxongDfiH7apYJeqGRj0X6szQFFVYFLqSzEZkVt73e3lffTq-9a0y_m0hyDrGI97f7WnmiqiMmVVRUU-78_2ttsPEmfgVwplbmRzdHJlYW0KZW5kb2JqCjkgMCBvYmoKPDwvRmlsdGVyL0ZsYXRlRGVjb2RlL0xlbmd0aCA3MDc-PnN0cmVhbQp4nKWW227aQBCG7_0Uc5lIybA7a--Bu8VZKK1jqO2gXESqlFSpVBVFaaT2ifqeHRsbcIrAqCAQK4ZvZv6dA6_RpIqUBmUdElRfo1BFnyMB9fPnt_ZDMYteazvZnCRIa1A4iKVEwT9aR6OpBAPVc3QxHfO3I-FGJCi-rL6fx7PodMezDS_cL8dgSY8kWWudlkeZJz3ZBEUMHLeKOz-68bMLGoQeJ2Jw6MKgIaBYoNsiacOc59VQTIyUgIolkusgbsMob33IoAofF_lRWJ-kOEMDT-t9H_V7w6ZW29RnWYCbAMviLkz8FQeAZ_gQstbyhA9fVWPgRxUydEqSihNtzlOFBCrZqSJFgzWJEBKuYemLeXkeTljU7h1uWvg8DedxdILK9YvIP75wsnW1CiUSGlxDfO9GHSRC9cFnoYSbeQmhXPo_fEsl-uPX9E-oiUVr-mC-9fH_wru4D-BBWDm8idooCVXSxyxxjsBFuvJ5Bemo-yTPlvYAu66dUHH23F6ruhPKh4uJL9KQLXL_cHlW7MTT0_QH1-LxbTyomSiuZ9-GcKKZFpMyFCufzhc5XxvH7yHzEPKqCDM_vHNJu3rOsLOT2bVzTTmLZN41ze3dXseggt_MUIZQOlhHZAVq1Z1_RG87k0RismfRHPcNyDkUemfRnnsmupnjO5PNuTE5nlC7aBTa7aKR1KST-zTcQy3wnAUGH4qwGFxmjtDowzItivksHJ_d-4LLuN6pbd12pE3FOt6BwyOyGCuQSuzNlDYi3kuhyJsy8tngwIRrUuQNmuymsdy0aTp4BvOGrQMilSDZfscUYcqa8xw-OjcaynWNk5pz09w5dWs_rWF0_yzh5gVOx0AS5aE_Ab8UcjnxrkrLVfxlMlhrLifO6aA0n2aDpekwJkZp3mHiq4GiWEKreprQME3a6z3knFftiF_bAP4COGUU4gplbmRzdHJlYW0KZW5kb2JqCjEwIDAgb2JqCjw8L0tpZHNbMTEgMCBSXS9UeXBlL1BhZ2VzL0NvdW50IDE-PgplbmRvYmoKMTEgMCBvYmoKPDwvQ29udGVudHNbMiAwIFIgOSAwIFIgMyAwIFJdL1R5cGUvUGFnZS9SZXNvdXJjZXM8PC9Gb250PDwvRjEgNCAwIFIvRjIgNSAwIFI-Pi9YT2JqZWN0PDwvWGYxIDYgMCBSL1hmMiA4IDAgUj4-Pj4vUGFyZW50IDEwIDAgUi9NZWRpYUJveFswIDAgMjgwLjYzIDQyNS4yXT4-CmVuZG9iagoxMiAwIG9iago8PC9UeXBlL0NhdGFsb2cvUGFnZXMgMTAgMCBSPj4KZW5kb2JqCnhyZWYKMCAxMwowMDAwMDAwMDAwIDY1NTM1IGYgCjAwMDAwMDAwMTUgMDAwMDAgbiAKMDAwMDAwMDI0NCAwMDAwMCBuIAowMDAwMDAwMzU4IDAwMDAwIG4gCjAwMDAwMDA0NDAgMDAwMDAgbiAKMDAwMDAwMDUzMyAwMDAwMCBuIAowMDAwMDAwNjIzIDAwMDAwIG4gCjAwMDAwMDEwOTIgMDAwMDAgbiAKMDAwMDAwMTE4MCAwMDAwMCBuIAowMDAwMDAxNjE4IDAwMDAwIG4gCjAwMDAwMDIzOTIgMDAwMDAgbiAKMDAwMDAwMjQ0NSAwMDAwMCBuIAowMDAwMDAyNjE4IDAwMDAwIG4gCnRyYWlsZXIKPDwvSW5mbyAxIDAgUi9JRCBbPDE0MDYwOGMwMmNlNWMwZDk1MTdmYzYxNmVkZDJmYjUxPjxlMGZlZjM3OGI1YzhmMDhkZGFlMzY3N2Q2Y2NjNmQ0ZT5dL1Jvb3QgMTIgMCBSL1NpemUgMTM-PgolaVRleHQtNS41LjcKc3RhcnR4cmVmCjI2NjUKJSVFT0YK`
	base64String := `JVBERi0xLjQKJeLjz9MKMSAwIG9iago8PC9Nb2REYXRlKEQ6MjAyNDA5MTgwODE4NDArMDInMDAnKS9DcmVhdGlvbkRhdGUoRDoyMDI0MDkxODA4MTg0MCswMicwMCcpL1Byb2R1Y2VyKGlUZXh0riA1LjUuNyCpMjAwMC0yMDE1IGlUZXh0IEdyb3VwIE5WIFwoQUdQTC12ZXJzaW9uXCk7IG1vZGlmaWVkIHVzaW5nIGlUZXh0riA1LjUuNyCpMjAwMC0yMDE1IGlUZXh0IEdyb3VwIE5WIFwoQUdQTC12ZXJzaW9uXCkpPj4KZW5kb2JqCjIgMCBvYmoKPDwvRmlsdGVyL0ZsYXRlRGVjb2RlL0xlbmd0aCA0OD4-c3RyZWFtCnicK-TiKlQw0LNUMABBIG1oomdgbGhqYKJgZKhnZGZgYGCmkJzLVcgVCMQAzzkJDwplbmRzdHJlYW0KZW5kb2JqCjMgMCBvYmoKPDwvRmlsdGVyL0ZsYXRlRGVjb2RlL0xlbmd0aCAxNj4-c3RyZWFtCnicUwjkKuTiCgRDAA4gAhIKZW5kc3RyZWFtCmVuZG9iago0IDAgb2JqCjw8L1N1YnR5cGUvVHlwZTEvVHlwZS9Gb250L0Jhc2VGb250L0hlbHZldGljYS1Cb2xkL0VuY29kaW5nL1dpbkFuc2lFbmNvZGluZz4-CmVuZG9iago1IDAgb2JqCjw8L1N1YnR5cGUvVHlwZTEvVHlwZS9Gb250L0Jhc2VGb250L1RpbWVzLVJvbWFuL0VuY29kaW5nL1dpbkFuc2lFbmNvZGluZz4-CmVuZG9iago2IDAgb2JqCjw8L1N1YnR5cGUvRm9ybS9GaWx0ZXIvRmxhdGVEZWNvZGUvVHlwZS9YT2JqZWN0L01hdHJpeFsxIDAgMCAxIDAgMF0vRm9ybVR5cGUgMS9SZXNvdXJjZXM8PC9Gb250PDwvRjMgNyAwIFI-Pj4-L0JCb3hbMCAwIDE3NS4zNSA2OS42Nl0vTGVuZ3RoIDI5MD4-c3RyZWFtCniclVIxUgQxDOvzipTQGCuJ46Rl5ihp2IIPcAUzNPy_IHC7uczaDa1GkSJLHDvVGhMhVo7fHyET5IaBWHawUjYYQGKZyFR39E9pRzt1OTulQTDM1KkY0TyszD_LNFpApWYkS_5VOmuWSsmCnbL9vQyrs5MoqXMnODmrHDlXZj9OeidqomLdtc6gy_vG1Bw0Tf9FoenstIzQN7Cn49ALs6tzVDAfV12L5uxkBTdvPwBTcmYBscUCbX7i_l2MtXRDTepVizEiGxhjMTYxcjvKXRUKZr0ruYiXWdjNPGaTHbKSIzyGo4Zaq9cmFK6d5nniRUNlfoL_iV3D8xaeXnJscbsGDJwjolRSHcaDuH2Fh9d3weUtNfTO_Lh9hssWfgD_mNx0CmVuZHN0cmVhbQplbmRvYmoKNyAwIG9iago8PC9TdWJ0eXBlL1R5cGUxL1R5cGUvRm9udC9CYXNlRm9udC9IZWx2ZXRpY2EvRW5jb2RpbmcvV2luQW5zaUVuY29kaW5nPj4KZW5kb2JqCjggMCBvYmoKPDwvU3VidHlwZS9Gb3JtL0ZpbHRlci9GbGF0ZURlY29kZS9UeXBlL1hPYmplY3QvTWF0cml4WzEgMCAwIDEgMCAwXS9Gb3JtVHlwZSAxL1Jlc291cmNlczw8L0ZvbnQ8PC9GMyA3IDAgUj4-Pj4vQkJveFswIDAgMTkyLjczIDc3LjE2XS9MZW5ndGggMjYwPj5zdHJlYW0KeJxtkztOBjEMhPucIiU0gx-xY7dIcIK9AhRINNy_IAVa6RejLSxNZibJp6zMRuZUWMzciPnzMeyM_-qf5th9ayoQ5tVTQfUzm_SYIRc7yUYW8bthG_F7wJi-BLaZHnAl_SHwZAwC7JixwW6VhmbUzuwmNdshrH4Hgup1-x_6K5BO9BYko9DBaXZxmioOS_ocivNUDbAbqxbYDdSCI1I7Abrgciceq7wQxRJLIXzB78Rj1SokBRIKNZaIwlaWyOOjCLNA9z6Po_gvdQKUSDmaPRutxKLUW7EYkc_xeo2Xd581r8-hU86nMwW7jvG4r-_xVEtFVJZ0VW-X5-trvF3jFzae4HkKZW5kc3RyZWFtCmVuZG9iago5IDAgb2JqCjw8L0ZpbHRlci9GbGF0ZURlY29kZS9MZW5ndGggNzM1Pj5zdHJlYW0KeJyVlV1T2kAUhu_zK86lzuhhP7LZXe4iLpaWAg2p44UznaDR6oAoWPuL-j97EoIQZWAzDAw7HJ59z5vz8RKcpYGMQBqLAtLbwKXBj4BB8VrcV1-Si-CliOPliQM3GpmFkHNk9KdZ0Opy0JDeBUfdNv3aYrYlmAiP08dmPIM2WvNMyXNXozYYEbW4MMZYE3kztUFlgCTKcI2MSuQg7rirgUt9SUYhC3eBNokCM21uvKUxjVqACBnad6RYMXsDb10hCgUy5CjsGmJXjOUsy6eQ5o_zp72wOklShhpuZtt3FJ8lW1TPo5NNpznc5vC8-JNPshMSgA3uYLzw8sAdcZq2AVLXRyu5kKGKdDNPBEPJ155wVkK1YozDKYzipDduhmMGI_sB103iQcc140QKpa2XUDyZU6pFfTPJlPCuIHrqWu4kQvol7rsxnPfG4Maj-F98AmOM9z-kT1KVQaPr4HMHbfiavT3cNlW5A0YNwxVrqEmgVHXM7zk8zak1Vetumr0W3ycclq-LPC8OjXV-vqAoF5eSm64Pl3GfjL0-OouTjusPB_H1caMEBI1YXZ9uw8my7dU9IiwG5IpwoHvmk2W-eMtuHuZP-RKes0UG0wzyJ3LlPvNvVRHZYrDQZQezqwaZtAaF_tAn339uNQlK-EsMqQVyC7NAGIaRXJ-nwXITojiqrYjyuB0grEUWbSKqcy0kKgf3JmR1LkP2J1RtI4nmfRtxsdkdMHbJZa_TG0LsEjf0LjMrUEe7bRomvQs38CUJHhaLt6rbNWlVsZYWpb8ig6EELtnWGKkU0SJyCaXbo0LvewtjtkyRVqbaDGBeAkcd77FLK7UQJKRCYeodk7gueU6jd-_wKCmnBY5HlFtEnVO09s0MWld3HM7ncFiD4Mh3bf03iVROtJ4648vw15m311ROlNNOa75deFuzxugQuf6A4SfKzxQj0MiaJ8LPk-rx7rqctmuL3u8C_gPYRCTTCmVuZHN0cmVhbQplbmRvYmoKMTAgMCBvYmoKPDwvS2lkc1sxMSAwIFJdL1R5cGUvUGFnZXMvQ291bnQgMT4-CmVuZG9iagoxMSAwIG9iago8PC9Db250ZW50c1syIDAgUiA5IDAgUiAzIDAgUl0vVHlwZS9QYWdlL1Jlc291cmNlczw8L0ZvbnQ8PC9GMSA0IDAgUi9GMiA1IDAgUj4-L1hPYmplY3Q8PC9YZjEgNiAwIFIvWGYyIDggMCBSPj4-Pi9QYXJlbnQgMTAgMCBSL01lZGlhQm94WzAgMCAyODAuNjMgNDI1LjJdPj4KZW5kb2JqCjEyIDAgb2JqCjw8L1R5cGUvQ2F0YWxvZy9QYWdlcyAxMCAwIFI-PgplbmRvYmoKeHJlZgowIDEzCjAwMDAwMDAwMDAgNjU1MzUgZiAKMDAwMDAwMDAxNSAwMDAwMCBuIAowMDAwMDAwMjQ0IDAwMDAwIG4gCjAwMDAwMDAzNTggMDAwMDAgbiAKMDAwMDAwMDQ0MCAwMDAwMCBuIAowMDAwMDAwNTMzIDAwMDAwIG4gCjAwMDAwMDA2MjMgMDAwMDAgbiAKMDAwMDAwMTA5MiAwMDAwMCBuIAowMDAwMDAxMTgwIDAwMDAwIG4gCjAwMDAwMDE2MTkgMDAwMDAgbiAKMDAwMDAwMjQyMSAwMDAwMCBuIAowMDAwMDAyNDc0IDAwMDAwIG4gCjAwMDAwMDI2NDcgMDAwMDAgbiAKdHJhaWxlcgo8PC9JbmZvIDEgMCBSL0lEIFs8MGEwODA1MDgzMWY0ZGQ1YTcwYmYxMzE2ZmIzODUxZWM-PDhiYWRiNTc2MmY5ODFmYWZkZjU5ODgwZDI0NjU4MTRjPl0vUm9vdCAxMiAwIFIvU2l6ZSAxMz4-CiVpVGV4dC01LjUuNwpzdGFydHhyZWYKMjY5NAolJUVPRgo*`
	// 替换所有的 '-' 为 '+'
	base64String = strings.ReplaceAll(base64String, "-", "+")
	// 替换所有的 '_' 为 '/'
	base64String = strings.ReplaceAll(base64String, "_", "/")

	base64String = strings.ReplaceAll(base64String, "*", "=")

	// Base64 解码
	decodedBytes, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		fmt.Println("Base64 解码失败:", err)
		return
	}

	// cleanBase64Str := ""
	// for _, char := range base64Str {
	// 	if (char >= 'A' && char <= 'Z') || (char >= 'a' && char <= 'z') || (char >= '0' && char <= '9') || char == '+' || char == '/' || char == '=' {
	// 		cleanBase64Str += string(char)
	// 	}
	// }
	// for len(cleanBase64Str)%4 != 0 {
	// 	cleanBase64Str += "="
	// }
	// fmt.Println(cleanBase64Str)

	// Decode the base64 string
	// decodedBytes, err := base64.StdEncoding.DecodeString(cleanBase64Str)
	// if err != nil {
	// 	fmt.Println("Error decoding base64:", err)
	// 	return
	// }
	fmt.Println(string(decodedBytes))
	// Write the decoded bytes to a PDF file
	// err = os.WriteFile("test_golang.pdf", decodedBytes, 0644)
	// if err != nil {
	// 	fmt.Println("Error writing file:", err)
	// 	return
	// }

	// fmt.Println("PDF file created successfully")
}
