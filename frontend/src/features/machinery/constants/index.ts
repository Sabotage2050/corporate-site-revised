import type { MachineCategory } from '../types'

   import lathe1 from '@/assets/machinery/lathe_machine/lathe1.jpg';
   import lathe2 from '@/assets/machinery/lathe_machine/lathe2.jpg';
   import lathe3 from '@/assets/machinery/lathe_machine/lathe3.jpg';
   import lathe4 from '@/assets/machinery/lathe_machine/lathe4.jpg';
   import lathe5 from '@/assets/machinery/lathe_machine/lathe5.jpg';

   import press1 from '@/assets/machinery/press_machine/press1.jpg';

   import drill1 from '@/assets/machinery/radial_drilling_machine/drill1.jpg';
   import drill2 from '@/assets/machinery/radial_drilling_machine/drill2.jpg';

   import mill1 from '@/assets/machinery/milling_machine/mill1.jpg';

   export const MACHINERY_CATEGORIES: MachineCategory[] = [
     {
       title: 'Lathe Machine',
       images: [
         { src: lathe1, alt: 'Lathe Machine 1', model: 'LM-001' },
         { src: lathe2, alt: 'Lathe Machine 2', model: 'LM-002' },
         { src: lathe3, alt: 'Lathe Machine 3', model: 'LM-003' },
         { src: lathe4, alt: 'Lathe Machine 4', model: 'LM-004' },
         { src: lathe5, alt: 'Lathe Machine 5', model: 'LM-005' },
       ]
     },
     {
       title: 'Press Machine',
       images: [
         { src: press1, alt: 'Press Machine 1', model: 'PM-001' },
       ]
     },
     {
       title: 'Radial Drilling Machine',
       images: [
         { src: drill1, alt: 'Radial Drilling Machine 1', model: 'RD-001' },
         { src: drill2, alt: 'Radial Drilling Machine 2', model: 'RD-002' },
       ]
     },
     {
       title: 'Milling Machine',
       images: [
         { src: mill1, alt: 'Milling Machine 1', model: 'MM-001' },
       ]
     }
   ];