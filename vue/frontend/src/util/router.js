import Vue from 'vue';
import Router from 'vue-router';
import Home from './../views/Home.vue';
// import DisplayInfo from './../views/DisplayInfo.vue';

Vue.use(Router);

export default new Router({
   // mode: 'history',
    routes: [
        {
            path: '/',
            name: 'home',
            component: Home
        },
        {
            path: '/displayInfo',
            name: 'displayInfo',
            // component: DisplayInfo
            component: () => import("./../views/DisplayInfo")
            // route level code-splitting
            // this generates a separate chunk (about.[hash].js) for this route
            // which is lazy-loaded when the route is visited.
            //component: () => import(/* webpackChunkName: "about" */ './views/Author.vue')
        },
        {
            path:'/authorList/:name',
            name: 'itemAuthor',
            component: ()=>import("../views/ItemAuthor")
        },
        {
            path: "/author/:id",
            name: "author",
            // component: About
            component: () => import('./../views/Author.vue')
        },
        {
            path:"/institution/:id",
            name:"institution",
            // component: Institution
            component: () => import('./../views/institution'),

        },
        {
            path:"/bioentity/:id",
            name:'bioentity',
            component: ()=>import('../views/Bioentity'),
        }



        /////////////////////////?/////////////////////
        // {
        //     path: "/medic",
        //     name: "medic",
        //     component: () => import("./../views/transplant/DisplayMedic.vue")
        // }
    ]


})

