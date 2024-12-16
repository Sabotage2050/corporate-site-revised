import { createRouter, createWebHistory } from 'vue-router'
import Home from '@/pages/Home.vue'
import About from '@/pages/About.vue'
import Machinery from '@/pages/Machinery.vue'
import Contact from '@/pages/Contact.vue'
import Result from '@/pages/Result.vue'
import StockListTop from '@/pages/stocklist/StockListTop.vue'
import ListForklift from '@/pages/stocklist/ListForklift.vue'
import ForkliftDetail from '@/pages/stocklist/ForkliftDetail.vue'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      name: 'home',
      component: Home
    },
    {
      path: '/about',
      name: 'about',
      component: About
    },
    {
      path: '/machinery',
      name: 'machinery',
      component: Machinery
    },
    {
      path: '/contact',
      name: 'contact',
      component: Contact
    },
    {
      path: '/contact/result',
      name: 'contact-result',
      component: Result
    },
    {
      path: '/stocklist',
      name: 'stocklist',
      component: StockListTop
    },
    {
      path: '/stocklist/forklift/:type',
      name: 'list-forklift',
      component: ListForklift,
      props: true
    },
    {
      path: '/stocklist/forklift/:type/:model/:serial',
      name: 'forklift-detail',
      component: ForkliftDetail,
      props: true
    }
  ]
})

export default router