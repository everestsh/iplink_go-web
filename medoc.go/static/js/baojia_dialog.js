// JavaScript Document
    function baojia_dialog(obj){
        var $dialog = $(
            "<div id='baojia_dialog'>"+
            "   <em></em><span></span>"+
            "   <div id='dialog'>"+
            "       <p id='dialog_title'>询问报价、售前或售后问题，请联系我们：</p>"+
            "       <p>公司座机：010-62416824(9:00-5:30)</p>"+
            "       <p>座机2&nbsp;&nbsp;&nbsp;：010-82168683(9:00-5:30)</p>"+
            "       <p>手机&nbsp;&nbsp;&nbsp;&nbsp;：152-1082-2019</p>"+
            "       <p>工程师1 ：138-1180-9455</p>"+
            "       <p>公司传真：010-62416834</p>"+
            "       <p>公司邮箱：sales@microembedded.com</p>"+
            "   </div>"+
            "</div>");
        var Top = 0;
        var Left = 0;
        $(obj).mouseenter(function(){
           var _top = $(this).offset().top;
           var _left = $(this).offset().left;
           Top = _top;
           Left = _left;
           $(this).append($dialog); 
           $dialog.css({"top":_top,"left":_left+120,"z-index":10});
           $dialog.stop().animate({
            left:_left+139,
            opacity:1
           });
           
           
            $(document).ready(function(){
                $(obj).bind('click', function(){
                    window.location.href = "../contact_us/contact_us.html";
                });
                $("#baojia_dialog").bind('click', function(){
                    return false;
                });
            });
        });
        
        $(obj).mouseleave(function(){
           $dialog.stop().animate({
               left:Left+120,
               opacity:0
           },function(){
               $dialog.remove();
           });
        });

    }
    baojia_dialog($(".xunwenbaojia"));