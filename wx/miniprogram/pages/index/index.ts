// index.ts
// 获取应用实例
const app = getApp<IAppOption>();

Page({
  isPageShowing: true,
  data: {
    setting: {
      skew: 0,
      rotate: 0,
      showLocation: true,
      showScale: true,
      subKey: "",
      layerStyle: 1,
      enableZoom: true,
      enableScroll: true,
      enableRotate: false,
      showCompass: false,
      enable3D: false,
      enableOverlooking: false,
      enableSatellite: false,
      enableTraffic: false,
    },
    location: {
      latitude: 23.09994,
      longitude: 113.32452,
    },
    scale: 12,
    markers: [
      {
        iconPath: "/resources/car.png",
        id: 0,
        latitude: 23.09994,
        longitude: 113.32452,
        width: 50,
        height: 50,
      },
      {
        iconPath: "/resources/car.png",
        id: 1,
        latitude: 23.09994,
        longitude: 114.32452,
        width: 50,
        height: 50,
      },
    ],
  },
  onLoad() {},
  onMyLocationTap() {
    wx.getLocation({
      type: "gcj02",
      success: (res) => {
        console.log({ res });
        this.setData({
          location: {
            latitude: res.latitude,
            longitude: res.longitude,
          },
        });
      },
      fail: () => {
        wx.showToast({
          icon: "none",
          title: "请前往授权页",
        });
      },
    });
  },
  onHide() {
    this.isPageShowing = false;
  },
  moveCars() {
    const map = wx.createMapContext("map");
    const dest = {
      latitude: 23.09994,
      longitude: 113.32452,
    };
    const moveCars = () => {
      dest.latitude += 0.1;
      dest.longitude += 0.1;
      map.translateMarker({
        destination: {
          latitude: dest.latitude,
          longitude: dest.longitude,
        },
        markerId: 0,
        autoRotate: false,
        rotate: 0,
        duration: 5000,
        animationEnd: () => { 
          if(this.isPageShowing) {
            moveCars()
          }
        }
      });
    };
    moveCars()
  },
  onScanClicked() {
    wx.scanCode({
      success: res => {
        wx.navigateTo({
          url: '/pages/register/register'
        })
      },
      fail: console.error
    })
  }
});
