(self.webpackChunk_N_E=self.webpackChunk_N_E||[]).push([[454,520,700,974],{1041:(e,r,t)=>{Promise.resolve().then(t.bind(t,6598))},6598:(e,r,t)=>{"use strict";t.d(r,{Label:()=>d});var n=t(5155),l=t(2115),i=t(3360),a=l.forwardRef((e,r)=>(0,n.jsx)(i.sG.label,{...e,ref:r,onMouseDown:r=>{var t;r.target.closest("button, input, select, textarea")||(null===(t=e.onMouseDown)||void 0===t||t.call(e,r),!r.defaultPrevented&&r.detail>1&&r.preventDefault())}}));a.displayName="Label";var o=t(1027),s=t(9602);let u=(0,o.F)("text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"),d=l.forwardRef((e,r)=>{let{className:t,...l}=e;return(0,n.jsx)(a,{ref:r,className:(0,s.cn)(u(),t),...l})});d.displayName=a.displayName},9602:(e,r,t)=>{"use strict";t.d(r,{cn:()=>i});var n=t(3463),l=t(9795);function i(){for(var e=arguments.length,r=Array(e),t=0;t<e;t++)r[t]=arguments[t];return(0,l.QP)((0,n.$)(r))}},8068:(e,r,t)=>{"use strict";t.d(r,{s:()=>a,t:()=>i});var n=t(2115);function l(e,r){if("function"==typeof e)return e(r);null!=e&&(e.current=r)}function i(...e){return r=>{let t=!1,n=e.map(e=>{let n=l(e,r);return t||"function"!=typeof n||(t=!0),n});if(t)return()=>{for(let r=0;r<n.length;r++){let t=n[r];"function"==typeof t?t():l(e[r],null)}}}}function a(...e){return n.useCallback(i(...e),e)}},3360:(e,r,t)=>{"use strict";t.d(r,{sG:()=>a});var n=t(2115);t(7650);var l=t(2317),i=t(5155),a=["a","button","div","form","h2","h3","img","input","label","li","nav","ol","p","span","svg","ul"].reduce((e,r)=>{let t=n.forwardRef((e,t)=>{let{asChild:n,...a}=e,o=n?l.DX:r;return"undefined"!=typeof window&&(window[Symbol.for("radix-ui")]=!0),(0,i.jsx)(o,{...a,ref:t})});return t.displayName=`Primitive.${r}`,{...e,[r]:t}},{})},2317:(e,r,t)=>{"use strict";t.d(r,{DX:()=>a});var n=t(2115),l=t(8068),i=t(5155),a=n.forwardRef((e,r)=>{let{children:t,...l}=e,a=n.Children.toArray(t),s=a.find(u);if(s){let e=s.props.children,t=a.map(r=>r!==s?r:n.Children.count(e)>1?n.Children.only(null):n.isValidElement(e)?e.props.children:null);return(0,i.jsx)(o,{...l,ref:r,children:n.isValidElement(e)?n.cloneElement(e,void 0,t):null})}return(0,i.jsx)(o,{...l,ref:r,children:t})});a.displayName="Slot";var o=n.forwardRef((e,r)=>{let{children:t,...i}=e;if(n.isValidElement(t)){let e=function(e){let r=Object.getOwnPropertyDescriptor(e.props,"ref")?.get,t=r&&"isReactWarning"in r&&r.isReactWarning;return t?e.ref:(t=(r=Object.getOwnPropertyDescriptor(e,"ref")?.get)&&"isReactWarning"in r&&r.isReactWarning)?e.props.ref:e.props.ref||e.ref}(t);return n.cloneElement(t,{...function(e,r){let t={...r};for(let n in r){let l=e[n],i=r[n];/^on[A-Z]/.test(n)?l&&i?t[n]=(...e)=>{i(...e),l(...e)}:l&&(t[n]=l):"style"===n?t[n]={...l,...i}:"className"===n&&(t[n]=[l,i].filter(Boolean).join(" "))}return{...e,...t}}(i,t.props),ref:r?(0,l.t)(r,e):e})}return n.Children.count(t)>1?n.Children.only(null):null});o.displayName="SlotClone";var s=({children:e})=>(0,i.jsx)(i.Fragment,{children:e});function u(e){return n.isValidElement(e)&&e.type===s}},1027:(e,r,t)=>{"use strict";t.d(r,{F:()=>a});var n=t(3463);let l=e=>"boolean"==typeof e?`${e}`:0===e?"0":e,i=n.$,a=(e,r)=>t=>{var n;if((null==r?void 0:r.variants)==null)return i(e,null==t?void 0:t.class,null==t?void 0:t.className);let{variants:a,defaultVariants:o}=r,s=Object.keys(a).map(e=>{let r=null==t?void 0:t[e],n=null==o?void 0:o[e];if(null===r)return null;let i=l(r)||l(n);return a[e][i]}),u=t&&Object.entries(t).reduce((e,r)=>{let[t,n]=r;return void 0===n||(e[t]=n),e},{});return i(e,s,null==r?void 0:null===(n=r.compoundVariants)||void 0===n?void 0:n.reduce((e,r)=>{let{class:t,className:n,...l}=r;return Object.entries(l).every(e=>{let[r,t]=e;return Array.isArray(t)?t.includes({...o,...u}[r]):({...o,...u})[r]===t})?[...e,t,n]:e},[]),null==t?void 0:t.class,null==t?void 0:t.className)}}},e=>{var r=r=>e(e.s=r);e.O(0,[181,441,517,358],()=>r(1041)),_N_E=e.O()}]);