"use strict";(self.webpackChunkmeteor=self.webpackChunkmeteor||[]).push([[840],{3905:function(e,t,n){n.d(t,{Zo:function(){return u},kt:function(){return m}});var r=n(7294);function i(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function o(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function a(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?o(Object(n),!0).forEach((function(t){i(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):o(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function s(e,t){if(null==e)return{};var n,r,i=function(e,t){if(null==e)return{};var n,r,i={},o=Object.keys(e);for(r=0;r<o.length;r++)n=o[r],t.indexOf(n)>=0||(i[n]=e[n]);return i}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(r=0;r<o.length;r++)n=o[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(i[n]=e[n])}return i}var l=r.createContext({}),c=function(e){var t=r.useContext(l),n=t;return e&&(n="function"==typeof e?e(t):a(a({},t),e)),n},u=function(e){var t=c(e.components);return r.createElement(l.Provider,{value:t},e.children)},p={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},d=r.forwardRef((function(e,t){var n=e.components,i=e.mdxType,o=e.originalType,l=e.parentName,u=s(e,["components","mdxType","originalType","parentName"]),d=c(n),m=i,g=d["".concat(l,".").concat(m)]||d[m]||p[m]||o;return n?r.createElement(g,a(a({ref:t},u),{},{components:n})):r.createElement(g,a({ref:t},u))}));function m(e,t){var n=arguments,i=t&&t.mdxType;if("string"==typeof e||i){var o=n.length,a=new Array(o);a[0]=d;var s={};for(var l in t)hasOwnProperty.call(t,l)&&(s[l]=t[l]);s.originalType=e,s.mdxType="string"==typeof e?e:i,a[1]=s;for(var c=2;c<o;c++)a[c]=n[c];return r.createElement.apply(null,a)}return r.createElement.apply(null,n)}d.displayName="MDXCreateElement"},4953:function(e,t,n){n.r(t),n.d(t,{frontMatter:function(){return s},contentTitle:function(){return l},metadata:function(){return c},toc:function(){return u},default:function(){return d}});var r=n(7462),i=n(3366),o=(n(7294),n(3905)),a=["components"],s={},l="Plugins",c={unversionedId:"guides/list_Plugins",id:"guides/list_Plugins",isDocsHomePage:!1,title:"Plugins",description:"Before getting started we expect you went through the prerequisites.",source:"@site/docs/guides/1_list_Plugins.md",sourceDirName:"guides",slug:"/guides/list_Plugins",permalink:"/meteor/docs/guides/list_Plugins",editUrl:"https://github.com/odpf/meteor/edit/master/docs/docs/guides/1_list_Plugins.md",tags:[],version:"current",sidebarPosition:1,frontMatter:{},sidebar:"docsSidebar",previous:{title:"Installation",permalink:"/meteor/docs/guides/installation"},next:{title:"Recipes - Creation and linting",permalink:"/meteor/docs/guides/manage_recipes"}},u=[{value:"Listing all the plugins",id:"listing-all-the-plugins",children:[]}],p={toc:u};function d(e){var t=e.components,n=(0,i.Z)(e,a);return(0,o.kt)("wrapper",(0,r.Z)({},p,n,{components:t,mdxType:"MDXLayout"}),(0,o.kt)("h1",{id:"plugins"},"Plugins"),(0,o.kt)("p",null,"Before getting started we expect you went through the ",(0,o.kt)("a",{parentName:"p",href:"/meteor/docs/guides/introduction#prerequisites"},"prerequisites"),"."),(0,o.kt)("p",null,"Meteor follows a plugin driven approach and hence includes basically three kinds of plugins for the metadata orchestration: extractors (source), processors, and sinks (destination).\nSome details on these 3 are:"),(0,o.kt)("ul",null,(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("p",{parentName:"li"},(0,o.kt)("strong",{parentName:"p"},"Extractors")," are the set of plugins that are source of our metadata and include databases, dashboards, users, etc.")),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("p",{parentName:"li"},(0,o.kt)("strong",{parentName:"p"},"Processors")," are the set of plugins that perform the enrichment or data processing for the metadata after extraction.")),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("p",{parentName:"li"},(0,o.kt)("strong",{parentName:"p"},"Sinks")," are the plugins that act as the destination of our metadata after extraction and processing."))),(0,o.kt)("p",null,"Read more about the concepts on each of these in ",(0,o.kt)("a",{parentName:"p",href:"/meteor/docs/concepts/overview"},"concepts"),".\nTo get more context on these plugins, it is recommended to try out the ",(0,o.kt)("inlineCode",{parentName:"p"},"list")," command to get the list of plugins of a specific type. Commands to list the plugins are mentioned below"),(0,o.kt)("h2",{id:"listing-all-the-plugins"},"Listing all the plugins"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-bash"},"# list all available extractors\n$ meteor list extractors\n\n# list all extractors with alias 'e'\n$ meteor list e\n\n# list available sinks\n$ meteor list sinks\n\n# list all sinks with alias 's'\n$ meteor list s\n\n# list all available processors\n$ meteor list processors\n\n# list all processors with alias 'p'\n$ meteor list p\n")))}d.isMDXComponent=!0}}]);