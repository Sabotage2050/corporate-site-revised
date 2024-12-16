import type { StockTopItem } from '../types'

import dieselTop from '@/assets/stocklist/dieselTop.jpg'
import gasolineTop from '@/assets/stocklist/gasolineTop.jpg'
import batteryTop from '@/assets/stocklist/batteryTop.jpg'
import shovelloadertop from '@/assets/stocklist/shovelloaderTop.jpg'

export const STOCK_ITEMS: StockTopItem[] = [
  {
    image: dieselTop,
    alt: 'Diesel Forklift',
    path: '/stocklist/forklift/diesel',
    label: 'Diesel'
  },
  {
    image: gasolineTop,
    alt: 'Gasoline Forklift',
    path: '/stocklist/forklift/gasoline',
    label: 'Gasoline'
  },
  {
    image: batteryTop,
    alt: 'Battery Forklift',
    path: '/stocklist/forklift/battery',
    label: 'Battery'
  },
  {
    image: shovelloadertop,
    alt: 'Shovelloader',
    path: '/stocklist/forklift/shovelloader',
    label: 'Shovelloader'
  }
]

// 各タイプごとの商品データ
export const FORKLIFT_ITEMS = {
  diesel: [
    { id: 'diesel1', image: '/src/assets/stocklist/diesel/model1.jpg', name: 'Diesel Model 1' },
    { id: 'diesel2', image: '/src/assets/stocklist/diesel/model2.jpg', name: 'Diesel Model 2' },
    // ... 他のディーゼルモデル
  ],
  gasoline: [
    { id: 'gas1', image: '/src/assets/stocklist/gasoline/model1.jpg', name: 'Gasoline Model 1' },
    // ... 他のガソリンモデル
  ],
  battery: [
    { id: 'bat1', image: '/src/assets/stocklist/battery/model1.jpg', name: 'Battery Model 1' },
    // ... 他のバッテリーモデル
  ],
  shovelloader: [
    { id: 'shv1', image: '/src/assets/stocklist/shovelloader/model1.jpg', name: 'Shovelloader Model 1' },
    // ... 他のショベルローダーモデル
  ]
}
