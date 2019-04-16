;(function(){
    var iconList=new Vue({
        el:'.icon',
        data:{
            iconSelect:[],
            iconsrc:'',
            text:'',
            id:'',
            list:[],
        },
        created:function(){
            var that=this;
            axios({
                method:'post',
                url:'/icon/iconlist.do',
                timeout: 1000,
                params:{},
            }).then(function(res){
                if(res.data.success==0){
                    that.iconSelect=res.data.data;
                    console.log(that.iconSelect);
                }
            });
            that.load();
        },
        methods:{
            //加载数据
            load:function(){
                var that=this;
                axios({
                    method:'post',
                    url:'/icon/selectlist.do',
                    timeout: 1000,
                    params:{},
                }).then(function(res){
                    if(res.data.success==0){
                        that.list=res.data.data;
                        console.log(that.iconSelect);
                    }
                });
            },
            selected:function(src){
                this.iconsrc=src;
            },
            add:function(){
                var item={};
                var that=this;
                item.name=this.text;
                item.icon=this.iconsrc;
                axios({
                    method:'post',
                    url:'/icon/iuItem.do',
                    timeout: 1000,
                    params:item,
                }).then(function(res){
                    if(res.data.success==0){
                        that.load();
                        that.text='';
                        that.iconsrc='';
                        that.id='';
                    }
                });

            },
            selectItem:function(index){
                var item=this.list[index];
                this.text=item.name;
                this.iconsrc=item.icon;
                this.id=item.id;
            },
            update:function(){
                var that=this;
                var item={id:this.id,name:this.text,icon:this.iconsrc};
                axios({
                    method:'post',
                    url:'/icon/iuItem.do',
                    timeout: 1000,
                    params:item,
                }).then(function(res){
                    if(res.data.success==0){
                        that.load();
                    }
                });
            },
        }
    });
})();