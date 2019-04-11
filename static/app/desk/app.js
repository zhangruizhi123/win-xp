;(function(){
    Vue.component('Toast',{
		template:`
		
		<div class="toast_contain" v-bind:style="{display:display}">
			<div class="mask"></div>
			<div class="text" v-bind:style="{color:color,backgroundColor:background}">{{text}}</div>
		</div>
		
		`,
		
		props:{
			text:{
				default:[]
            },
            display:{
                default:'block'
            },
            color:{default: '#fff'},
            background:{default:'#000'},
		}
		,
		data:function(){
			return{
				menudisplay:'none',
				menuLeft:'0px',
				menuTop:'0px',
			}
		},
		methods:{
            
		}
	});
})();