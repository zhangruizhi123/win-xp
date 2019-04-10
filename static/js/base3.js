;(function(){
	var win_xp_dlg =Vue.component('win-xp-dlg',{
		template:`
		<div class="winxp_dlg base_size" v-bind:style="{width:width,height:height,left:left,top:top,cursor:cursor}" v-if="!closed" @mousedown="downresize" @mouseleave="leaveresize" @mouseenter="enterresize" >
			<!--- 头 --->
			<div class="winxp_dlg_head" @mousedown="down" @mouseup="up" @mouseleave="leave" >
				<img class="winxp_dlg_head_img" v-bind:src="icon" >
				<div class="winxp_dlg_head_title">{{title}}</div>
				<div class="winxp_dlg_head_btngup" >
					<div class="btn min" v-on:click="min"></div>
					<div class="btn" :class="cls" v-on:click="max"></div>
					<div class="btn close" v-on:click="close"></div>
				</div>
			</div>
			<div class="winxp_dlg_content" @mousedown="bydown" >
				<slot></slot>
			</div>
			<!--
			<div class="winxp_dlg_footer">
				<slot name="footer" ></slot>
			</div>
			-->
		</div>
		`,
		props:{
			title:{
				type:String,
				default:'我的电脑1'
			},
			icon:{
				type:String,
				default:'img/computer.png'
			}
		}
		,
		data:function(){
			return {
			width:'500px',
			height:'500px',
			leftVal:20,
			topVal:20,
			left:'20px',
			top:'20px',
			closed:false,
			cls:'max',
			cursor:'default',
			}
		},
		created:function(){
			
		},
		methods:{
			/**最小化**/
			min:function(e){
				var that=this;
				if(this.cls!='max'){
					this.left='0px';
					this.top='0px';
					this.leftVal=0;
					this.topVal=0;
					this.width='500px';
					this.height='500px';
					this.cls='max';
				}
				this.$emit('min');
			},
			/***最大化**/
			max:function(e){
				if(this.cls=='max'){
					this.left='0px';
					this.top='0px';
					this.leftVal=0;
					this.topVal=0;
					this.width=(document.documentElement.clientWidth-6)+'px';
					this.height=(document.documentElement.clientHeight-41)+'px';
					this.cls='minimize';
				}else{
					this.left='0px';
					this.top='0px';
					this.leftVal=0;
					this.topVal=0;
					this.width='500px';
					this.height='500px';
					this.cls='max';
				}
				this.$emit('max');
			},
			/**关闭**/
			close:function(e){
				this.closed=true;
				this.$emit('close');
			},
			/**鼠标按下事件**/
			down:function(event){
				var that=this;
				var x=that.leftVal;
				var y=that.topVal;
				var event=event||window.event;
				var startx=event.clientX;
				var starty=event.clientY;
				var sb_bkx=startx-event.target.offsetLeft;
				var sb_bky=starty-event.target.offsetTop;
				var ww=document.documentElement.clientWidth;
				var wh = window.innerHeight;
  				document.onmousemove=function(ev){
  					  var event=ev||window.event;
					   var scrolltop=document.documentElement.scrollTop||document.body.scrollTop;
					   if (event.clientY < 0 || event.clientX < 0 || event.clientY > wh || event.clientX > ww) {
					    return false;
					   };
					   var endx=event.clientX-sb_bkx;
					   var endy=event.clientY-sb_bky;
					   that.left=(x+endx)+'px';
					   that.leftVal=(x+endx);
					   that.top=(y+endy)+'px';
					   that.topVal=(y+endy);
  				}
  				//阻止冒泡事件
				event.stopPropagation();
				return false;
			},
			bydown:function(event){
				//阻止冒泡事件
				event.stopPropagation();
			},
			/**鼠标释放**/
			up:function(event){
				document.onmousemove=null;
			},
			/***鼠标移出事件**/
			leave:function(){
				document.onmousemove=null;
				//this.cursor='s-resize';
			},
			/*鼠标移出调整大小**/
			leaveresize:function(event){
				this.cursor='default';
				
			},
			//鼠标移入事件
			enterresize:function(event){
				console.log("++++");
				var width=parseInt(event.target.style.width.replace('px',''));
				var height=parseInt(event.target.style.height.replace('px',''));
				var startx=event.clientX;
				var starty=event.clientY;
				var x=startx-this.leftVal;
				var y=starty-this.topVal;
				if(x<5){
					console.log("left...")
					this.cursor='w-resize';
				}else if(y<5){
					this.cursor='s-resize';
				}else if(x>width){
					this.cursor='w-resize';
				}else if(y>height){
					this.cursor='s-resize';
				}
				console.log(x);
			},
			
			downresize:function(event){
				var that=this;
				var ww=document.documentElement.clientWidth;
				var wh = window.innerHeight;
				document.onmousemove=null;
				document.onmousemove=function(ev){
  					  var event=ev||window.event;
					   if (event.clientY < 0 || event.clientX < 0 || event.clientY > wh || event.clientX > ww) {
					    return false;
					   };
					  var endx=event.clientX;
					  var endy=event.clientY;
					  console.log(endy);
  				}
			},
		}
	});
	
	
})();
