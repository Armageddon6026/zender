<script setup lang="ts">
import { useUserStore } from '@/stores/UserStore'
import { Enviroment } from '@/common/const'
import LogosUnocss from '@/assets/svg/LogosUnocss.vue'
import DeviconGrpc from '@/assets/svg/DeviconGrpc.vue'

const ourUserStore = useUserStore()
</script>

<template>
  <div c="white">
    <h1 text="center">
      Wellcome {{ ourUserStore.UserInfo ? ourUserStore.UserInfo.name : 'Unknown' }}
    </h1>
    <div overflow="y-scroll x-hidden" h-75vh>
      <div zender-form bg="gray-8" h="40" w="190">
        <div flex="~">
          <i i="ic-round-linear-scale" m="t5 l3" text="size-30px teal"></i>
          <div m="t4 l4 b2" text=" size-3xl">Introduction</div>
        </div>
        <div m="l4 r4">
          Project Zender is a monitoring system written in Go language and Vue.js, providing simple
          and user-friendly real-time data and status monitoring across multiple servers.
        </div>
        <div m="t1 l4 r4">
          Front-End: Supported by <i i="fa6-brands-vuejs" text="emerald">XX</i>VueJs &
          <LogosUnocss /> Unocss。
        </div>
        <div m="l4 r4">
          Back-End: Supported by<i i="fa6-brands-golang" text="cyan">XX</i> Golang Gin &
          <DeviconGrpc size="20px" /> GRPC 。
        </div>
      </div>
      <div zender-form bg="gray-8" m="t5" h="40" w="190">
        <div flex="~">
          <i i="ic-twotone-dashboard-customize" m="t5 l3" text="size-30px lightBlue">x</i>
          <div m-4 text=" size-3xl">Applications Version</div>
        </div>
        <div flex="~" text="size-2xl">
          <div flex="~">
            <i i="ic-baseline-computer" m="l5 r1" text-orange>x</i>
            <div>Front-End</div>
            <span ml-2>: v{{ Enviroment.VERSION_FRONT_END }}</span>
          </div>
          <div flex="~">
            <i i="ic-round-view-compact-alt" m="l5 r1" text-emerald>x</i>
            <div>Back-End</div>
            <span ml-2>: v{{ Enviroment.VERSION_BACK_END }}</span>
          </div>
          <div flex="~">
            <i i="ic-twotone-shape-line" m="l5 r1" text-fuchsia>x</i>
            <div>SubSystem</div>
            <span ml-2>: v{{ Enviroment.VERSION_SUBSYSTEM }}</span>
          </div>
        </div>
        <button un-border="orange solid 1px rd-6px" m="4" h="8" w="40" bg="amber" text="size-18px">
          Change Log
        </button>
        <button
          un-border="gray solid 1px rd-6px"
          m="4"
          h="8"
          w="40"
          bg="gray-6"
          text="size-18px white"
        >
          <i i="fa6-brands-github">XX</i>
          GitHub
        </button>
      </div>
      <div bg="gray-8" m="t5" h="95" w="190" zender-form>
        <div flex="~">
          <i i="ic-outline-sticky-note-2" m="t5 l3" text="size-30px pink">x</i>
          <div m-4 text=" size-3xl">Licence</div>
        </div>
        <h3 m="t0" text="center">
          The MIT License (MIT) Copyright (c) 2024 Armageddon6026(HsuanYu,Liu)
        </h3>
        <div m="r5 l5" text="justify">
          Permission is hereby granted, free of charge, to any person obtaining a copy of this
          software and associated documentation files (the "Software"), to deal in the Software
          without restriction, including without limitation the rights to use, copy, modify, merge,
          publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons
          to whom the Software is furnished to do so, subject to the following conditions: The above
          copyright notice and this permission notice shall be included in all copies or substantial
          portions of the Software. <br />THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY
          KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
          FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
          COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN
          ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
          SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
        </div>
      </div>
    </div>
  </div>
</template>
