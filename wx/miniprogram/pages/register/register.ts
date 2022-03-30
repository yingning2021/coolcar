// pages/register/register.ts
Page({
    /**
     * 页面的初始数据
     */
    data: {
        genderIndex: 0,
        genders: ['未知', '男', '女', '其他'],
        hasLicImg: false,
        licIMgURL: "/resources/sedan.png" ,
    },
    onUploadLic() {
        wx.chooseImage({
            success: res => {
                if(res.tempFilePaths.length > 0) {
                    this.setData({
                        hasLicImg: true,
                        licIMgURL: res.tempFilePaths[0]
                    })
                }
            }
        })
    },
    onGenderChange(e: any) {
        this.setData({
            genderIndex: e.detail.value
        })
    }


})