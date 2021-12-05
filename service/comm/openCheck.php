<?php
// 获取毫秒时间戳
function getMillisecond() {
    list($s1,$s2)=explode(' ',microtime());
    return (float)sprintf('%.0f',(floatval($s1)+floatval($s2))*1000);
}

// 开放参数请求检查 type1检查评论
function openCheckErr($body,$type) {
  $timestamp = $body -> timestamp;
  $time = getMillisecond();
  $timeSub = $time - $timestamp;
  // 10秒非法
  if ( $timeSub < 0 || $timeSub > 10000 ) {
    return true;
  }
  // 评论参数非法
  if ( $type == 1 ) {
    $levelC = $body -> level;
    $nameC = $body -> name;
    $comsC = $body -> coms;
    if ( $levelC >2 ||  $levelC<1) {
      return true;
    }
    if ( $nameC == "" || $comsC == "") {
       return true;
    }
  }
  return false;
}
?>