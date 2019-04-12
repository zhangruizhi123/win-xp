;(function(){
	var win_xp_dlg =Vue.component('win-xp-dlg',{
		template:`
		<div >
		
		<div v-for="(item,index) in list"   class="winxp_dlg base_size" v-bind:ids="item.id" v-bind:style="{width:attrs[index].width,height:attrs[index].height,left:attrs[index].left,top:attrs[index].top,display:attrs[index].display,cursor:cursor}" v-if="!attrs[index].closed" @mousedown="downresize" @mouseleave="leaveresize" @mouseenter="enterresize" >
			<!--- 头 --->
			<div class="winxp_dlg_head" v-bind:ids="item.id" @mousedown="down" @mouseup="up" @mouseleave="leave" >
				<img class="winxp_dlg_head_img" v-bind:src="item.icon" >
				<div class="winxp_dlg_head_title">{{item.name}}</div>
				<div class="winxp_dlg_head_btngup" >
					<div class="btn min" v-on:click="min(item.id)"></div>
					<div class="btn" :class="cls" v-on:click="max(item.id)"></div>
					<div class="btn close" v-on:click="close(item.id)"></div>
				</div>
			</div>
			<div class="winxp_dlg_content" @mousedown="bydown" >
				<!-- <slot></slot> -->
				<iframe style="width:100%;height:100%;margin: 0px; padding: 0px;border:0px;" v-bind:src="item.url" ></iframe>
			</div>
			<!--
			<div class="winxp_dlg_footer">
				<slot name="footer" ></slot>
			</div>
			-->
		</div>
		
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
			},
			list:{
				default:[],
			},
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
			attrs:[],
			param:{name:"lis",age:22},
			}
		},
		
		created:function(){
			/*
			console.log("组件创建");
			console.log("++++++"+this.list.length);
			for(var i=0;i<this.list.length;i++){
				console.log(this.list[i].name)
			}
			*/
		},
		
		methods:{
			/**添加窗口事件**/
			additem:function(obj){
				var item={
					width:'500px',
					height:'500px',
					leftVal:20,
					topVal:20,
					left:'20px',
					top:'20px',
					closed:false,
					cls:'max',
					display:'block',
					id:obj.id
				};
				this.attrs.push(item)
			},
			showstate:function(id){
				var item={};
				for(var i=0;i<this.attrs.length;i++){
					if(this.attrs[i].id==id){
						item=this.attrs[i];
						break;
					}
				}
				
				if(item.display=='block'){
					item.display='none';
				}else{
					item.display='block';
				}
			},
			/**最小化**/
			min:function(id){
				var that=this;
				var item={};
				for(var i=0;i<this.attrs.length;i++){
					if(this.attrs[i].id==id){
						item=this.attrs[i];
						break;
					}
				}
				item.display='none';
				this.$emit('min',id);
			},
			/***最大化**/
			max:function(id){
				var item={};
				for(var i=0;i<this.attrs.length;i++){
					if(this.attrs[i].id==id){
						item=this.attrs[i];
						break;
					}
				}
				if(item.cls=='max'){
					item.left='0px';
					item.top='0px';
					item.leftVal=0;
					item.topVal=0;
					item.width=(document.documentElement.clientWidth-6)+'px';
					item.height=(document.documentElement.clientHeight-41)+'px';
					item.cls='minimize';
				}else{
					item.left='0px';
					item.top='0px';
					item.leftVal=0;
					item.topVal=0;
					item.width='500px';
					item.height='500px';
					item.cls='max';
				}
				this.$emit('max',id);
			},
			/**关闭**/
			close:function(id){
				var item={};
				for(var i=0;i<this.attrs.length;i++){
					if(this.attrs[i].id==id){
						item=this.attrs[i];
						this.attrs.splice(i,1);
						break;
					}
				}
				//item.closed=true;
				this.$emit('close',id);
			},
			/**鼠标按下事件**/
			down:function(event){
				var id=event.target.getAttribute("ids");
				
				var item={};
				for(var i=0;i<this.attrs.length;i++){
					if(this.attrs[i].id==id){
						item=this.attrs[i];
						break;
					}
				}
				var that=item;
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
				//this.cursor='default';
				
			},
			//鼠标移入事件
			enterresize:function(event){
				/*
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
				*/
			},
			
			downresize:function(event){
				/*
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
				
				*/
			},
		}
	});
	
	
})();
