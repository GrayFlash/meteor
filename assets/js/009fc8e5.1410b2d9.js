"use strict";(self.webpackChunkmeteor=self.webpackChunkmeteor||[]).push([[38],{3905:function(e,n,r){r.d(n,{Zo:function(){return l},kt:function(){return d}});var t=r(7294);function i(e,n,r){return n in e?Object.defineProperty(e,n,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[n]=r,e}function o(e,n){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var t=Object.getOwnPropertySymbols(e);n&&(t=t.filter((function(n){return Object.getOwnPropertyDescriptor(e,n).enumerable}))),r.push.apply(r,t)}return r}function a(e){for(var n=1;n<arguments.length;n++){var r=null!=arguments[n]?arguments[n]:{};n%2?o(Object(r),!0).forEach((function(n){i(e,n,r[n])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):o(Object(r)).forEach((function(n){Object.defineProperty(e,n,Object.getOwnPropertyDescriptor(r,n))}))}return e}function s(e,n){if(null==e)return{};var r,t,i=function(e,n){if(null==e)return{};var r,t,i={},o=Object.keys(e);for(t=0;t<o.length;t++)r=o[t],n.indexOf(r)>=0||(i[r]=e[r]);return i}(e,n);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(t=0;t<o.length;t++)r=o[t],n.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(i[r]=e[r])}return i}var c=t.createContext({}),p=function(e){var n=t.useContext(c),r=n;return e&&(r="function"==typeof e?e(n):a(a({},n),e)),r},l=function(e){var n=p(e.components);return t.createElement(c.Provider,{value:n},e.children)},u={inlineCode:"code",wrapper:function(e){var n=e.children;return t.createElement(t.Fragment,{},n)}},m=t.forwardRef((function(e,n){var r=e.components,i=e.mdxType,o=e.originalType,c=e.parentName,l=s(e,["components","mdxType","originalType","parentName"]),m=p(r),d=i,f=m["".concat(c,".").concat(d)]||m[d]||u[d]||o;return r?t.createElement(f,a(a({ref:n},l),{},{components:r})):t.createElement(f,a({ref:n},l))}));function d(e,n){var r=arguments,i=n&&n.mdxType;if("string"==typeof e||i){var o=r.length,a=new Array(o);a[0]=m;var s={};for(var c in n)hasOwnProperty.call(n,c)&&(s[c]=n[c]);s.originalType=e,s.mdxType="string"==typeof e?e:i,a[1]=s;for(var p=2;p<o;p++)a[p]=r[p];return t.createElement.apply(null,a)}return t.createElement.apply(null,r)}m.displayName="MDXCreateElement"},97:function(e,n,r){r.r(n),r.d(n,{frontMatter:function(){return s},contentTitle:function(){return c},metadata:function(){return p},toc:function(){return l},default:function(){return m}});var t=r(7462),i=r(3366),o=(r(7294),r(3905)),a=["components"],s={},c="Recipes - Creation and linting",p={unversionedId:"guides/manage_recipes",id:"guides/manage_recipes",isDocsHomePage:!1,title:"Recipes - Creation and linting",description:"A recipe is a set of instructions and configurations defined by user, and in Meteor they are used to define how a particular job will be performed.",source:"@site/docs/guides/2_manage_recipes.md",sourceDirName:"guides",slug:"/guides/manage_recipes",permalink:"/meteor/docs/guides/manage_recipes",editUrl:"https://github.com/odpf/meteor/edit/master/docs/docs/guides/2_manage_recipes.md",tags:[],version:"current",sidebarPosition:2,frontMatter:{},sidebar:"docsSidebar",previous:{title:"Plugins",permalink:"/meteor/docs/guides/list_Plugins"},next:{title:"Running Meteor",permalink:"/meteor/docs/guides/run_recipes"}},l=[{value:"Generating Sample recipe(s)",id:"generating-sample-recipes",children:[]},{value:"Linting Recipe(s)",id:"linting-recipes",children:[]}],u={toc:l};function m(e){var n=e.components,r=(0,i.Z)(e,a);return(0,o.kt)("wrapper",(0,t.Z)({},u,r,{components:n,mdxType:"MDXLayout"}),(0,o.kt)("h1",{id:"recipes---creation-and-linting"},"Recipes - Creation and linting"),(0,o.kt)("p",null,"A recipe is a set of instructions and configurations defined by user, and in Meteor they are used to define how a particular job will be performed.\nThus, for the entire set of orchestration all you will need to provide will be recipe","(","s",")"," for all the jobs you want meteor to do."),(0,o.kt)("p",null,"Read more about the concepts of Recipe ",(0,o.kt)("a",{parentName:"p",href:"/meteor/docs/concepts/recipe"},"here"),"."),(0,o.kt)("p",null,"A sample recipe can be generated using the commands mentioned ",(0,o.kt)("a",{parentName:"p",href:"#generating-sample-recipes"},"below"),".\nAfter making the necessary changes to the source, and sinks as per ypur local setup, you can validate the sample-recipe using steps mentioed ",(0,o.kt)("a",{parentName:"p",href:"#linting-recipes"},"here"),"."),(0,o.kt)("h2",{id:"generating-sample-recipes"},"Generating Sample recipe","(","s",")"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-bash"},"# generate a sample recipe\n# generate a recipe with a bigquery extractor and a console sink\n$ meteor gen recipe sample -e <name-of-extractor> -s <single-or-multiple-sinks> -p <name-of-processors>\n\n# command to generate recipe with multiple sinks\n$ meteor gen recipe sample -e bigquery -s columbus,kafka\n\n# for the tour you can use a single console sink\n# extracor(-e) as postgres, sink(-s) and enrich processor(-p)\n# save the generated recipe to a recipe.yaml\nmeteor gen recipe sample -e postgres -s console -p enrich > recipe.yaml\n")),(0,o.kt)("h2",{id:"linting-recipes"},"Linting Recipe","(","s",")"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-bash"},"# validate specified recipes.\n$ meteor lint recipe.yml\n")),(0,o.kt)("p",null,"More options for lint and gen commands can be found ",(0,o.kt)("a",{parentName:"p",href:"/meteor/docs/reference/commands"},"here"),"."))}m.isMDXComponent=!0}}]);