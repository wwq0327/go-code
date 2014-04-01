package main

import (
	"fmt"
)

func main() {
	str := `<td class="cur">
                <dl>
                    <dt>1号</dt>
                    <dd>
                        <a onmouseout="hidee(this)" onmouseover="show(this)" href="http://www.yyets.com/resource/28516" target="_blank">
                        <div class="floatSpan" style="display: none;"><span> 杀手信徒 The Following  S02E11<span></span></span></div>
                        <font class="fa1">杀手信徒</font>
                        <font class="fa1" style="color:white">S02E11</font>
                        </a>

                    </dd>
                  <dd class="even">
                        <a onmouseout="hidee(this)" onmouseover="show(this)" href="http://www.yyets.com/resource/30032" target="_blank">
                        <div class="floatSpan" style="display: none;"><span> 未来青年 The Tomorrow People  S01E18<span></span></span></div>
                        <font class="fa1">未来青年</font>
                        <font class="fa1" style="color:white">S01E18</font>
                        </a>

                    </dd>
                  <dd>
                        <a onmouseout="hidee(this)" onmouseover="show(this)" href="http://www.yyets.com/resource/27750" target="_blank">
                        <div class="floatSpan"><span> 皇家律师 Silk  S03E06<span></span></span></div>
                        <font class="fa1">皇家律师</font>
                        <font class="fa1" style="color:white">S03E06</font>
                        </a>

                    </dd>
                  <dd class="even">
                        <a onmouseout="hidee(this)" onmouseover="show(this)" href="http://www.yyets.com/resource/27621" target="_blank">
                        <div class="floatSpan"><span> 家族风云 Dallas  S03E06<span></span></span></div>
                        <font class="fa1">家族风云</font>
                        <font class="fa1" style="color:white">S03E06</font>
                        </a>

                    </dd>
                  <dd>
                        <a onmouseout="hidee(this)" onmouseover="show(this)" href="http://www.yyets.com/resource/29965" target="_blank">
                        <div class="floatSpan"><span> 智能缉凶 Intelligence  S01E13<span></span></span></div>
                        <font class="fa1">智能缉凶</font>
                        <font class="fa1" style="color:white">S01E13</font>
                        </a>

                    </dd>
                  <dd class="even">
                        <a onmouseout="hidee(this)" onmouseover="show(this)" href="http://www.yyets.com/resource/26154" target="_blank">
                        <div class="floatSpan"><span> 老爸老妈的浪漫史 How I Met Your Mother  S09E23<span></span></span></div>
                        <font class="fa1">老爸老妈的浪漫史</font>
                        <font class="fa1" style="color:white">S09E23</font>
                        </a>

                    </dd>
                  <dd>
                        <a onmouseout="hidee(this)" onmouseover="show(this)" href="http://www.yyets.com/resource/26154" target="_blank">
                        <div class="floatSpan"><span> 老爸老妈的浪漫史 How I Met Your Mother  S09E24<span></span></span></div>
                        <font class="fa1">老爸老妈的浪漫史</font>
                        <font class="fa1" style="color:white">S09E24</font>
                        </a>

                    </dd>
                  <dd class="even">
                        <a onmouseout="hidee(this)" onmouseover="show(this)" href="http://www.yyets.com/resource/11171" target="_blank">
                        <div class="floatSpan"><span> 我欲为人 Being Human US  S04E12<span></span></span></div>
                        <font class="fa1">我欲为人</font>
                        <font class="fa1" style="color:white">S04E12</font>
                        </a>

                    </dd>
                  <dd>
                        <a onmouseout="hidee(this)" onmouseover="show(this)" href="http://www.yyets.com/resource/10452" target="_blank">
                        <div class="floatSpan"><span> 射手 Archer  S05E10<span></span></span></div>
                        <font class="fa1">射手</font>
                        <font class="fa1" style="color:white">S05E10</font>
                        </a>

                    </dd>
                  <dd class="even">
                        <a onmouseout="hidee(this)" onmouseover="show(this)" href="http://www.yyets.com/resource/31346" target="_blank">
                        <div class="floatSpan"><span> 瑞克与莫蒂 Rick and Morty  S01E10<span></span></span></div>
                        <font class="fa1">瑞克与莫蒂</font>
                        <font class="fa1" style="color:white">S01E10</font>
                        </a>

                    </dd>
                  <dd>
                        <a onmouseout="hidee(this)" onmouseover="show(this)" href="http://www.yyets.com/resource/30659" target="_blank">
                        <div class="floatSpan"><span> 极品老妈 Mom  S01E21<span></span></span></div>
                        <font class="fa1">极品老妈</font>
                        <font class="fa1" style="color:white">S01E21</font>
                        </a>

                    </dd>
                  <dd class="even">
                        <a onmouseout="hidee(this)" onmouseover="show(this)" href="http://www.yyets.com/resource/31875" target="_blank">
                        <div class="floatSpan"><span> 独夫 The widower  S101E03<span></span></span></div>
                        <font class="fa1">独夫</font>
                        <font class="fa1" style="color:white">E03</font>
                        </a>

                    </dd>
                  <dd>
                        <a onmouseout="hidee(this)" onmouseover="show(this)" href="http://www.yyets.com/resource/29964" target="_blank">
                        <div class="floatSpan"><span> 罪恶黑名单 The Blacklist  S01E18<span></span></span></div>
                        <font class="fa1">罪恶黑名单</font>
                        <font class="fa1" style="color:white">S01E18</font>
                        </a>

                    </dd>
                  <dd class="even">
                        <a onmouseout="hidee(this)" onmouseover="show(this)" href="http://www.yyets.com/resource/31940" target="_blank">
                        <div class="floatSpan"><span> 损友的美好时代 Friends With Better Lives  S01E01<span></span></span></div>
                        <font class="fa1">损友的美好时代</font>
                        <font class="fa1" style="color:white">S01E01</font>
                        </a>

                    </dd>
                  <dd>
                        <a onmouseout="hidee(this)" onmouseover="show(this)" href="http://www.yyets.com/resource/11097" target="_blank">
                        <div class="floatSpan"><span> 识骨寻踪 Bones  S09E19<span></span></span></div>
                        <font class="fa1">识骨寻踪</font>
                        <font class="fa1" style="color:white">S09E19</font>
                        </a>

                    </dd>
                                  </dl>
            </td>
`

	fmt.Println(str)
}
