/*! For license information please see c4f5d8e4.b4a72cce.js.LICENSE.txt */
(self.webpackChunkmeteor=self.webpackChunkmeteor||[]).push([[195],{2579:function(e,t,a){"use strict";var n=a(4184),r=a.n(n),o=a(7294),i=function(e){var t,a=r()(e.className,{darkBackground:"dark"===e.background,highlightBackground:"highlight"===e.background,lightBackground:"light"===e.background,paddingAll:e.padding.indexOf("all")>=0,paddingBottom:e.padding.indexOf("bottom")>=0,paddingLeft:e.padding.indexOf("left")>=0,paddingRight:e.padding.indexOf("right")>=0,paddingTop:e.padding.indexOf("top")>=0});return t=e.wrapper?o.createElement("div",{className:"container"},e.children):e.children,o.createElement("div",{className:a,id:e.id},t)};i.defaultProps={background:null,padding:[],wrapper:!0},t.Z=i},9260:function(e,t,a){"use strict";var n=a(1721),r=a(4184),o=a.n(r),i=a(7294),s=function(e){function t(){return e.apply(this,arguments)||this}(0,n.Z)(t,e);var a=t.prototype;return a.renderBlock=function(e){var t=Object.assign({},{imageAlign:"left"},e),a=o()("blockElement",this.props.className,{alignCenter:"center"===this.props.align,alignRight:"right"===this.props.align,fourByGridBlock:"fourColumn"===this.props.layout,threeByGridBlock:"threeColumn"===this.props.layout,twoByGridBlock:"twoColumn"===this.props.layout});return i.createElement("div",{className:a,key:t.title},i.createElement("div",{className:"blockContent"},this.renderBlockTitle(t.title),t.content))},a.renderBlockTitle=function(e){return e?i.createElement("h2",null,e):null},a.render=function(){return i.createElement("div",{className:"gridBlock"},this.props.contents.map(this.renderBlock,this))},t}(i.Component);s.defaultProps={align:"left",contents:[],layout:"twoColumn"},t.Z=s},2841:function(e,t,a){"use strict";a.r(t),a.d(t,{default:function(){return u}});var n=a(7294),r=a(6698),o=a(6010),i=a(2263),s=a(2579),l=a(9260),c=a(4996),d=function(){var e=(0,i.Z)().siteConfig;return n.createElement("div",{className:"homeHero"},n.createElement("div",{className:"logo"},n.createElement("img",{src:(0,c.Z)("img/pattern.svg")})),n.createElement("div",{className:"container banner"},n.createElement("div",{className:"row"},n.createElement("div",{className:(0,o.Z)("col col--5")},n.createElement("div",{className:"homeTitle"},e.tagline),n.createElement("small",{className:"homeSubTitle"},"Meteor is an open source plugin-driven metadata collection framework to extract data from different sources and sink to any data catalog or store."),n.createElement("a",{className:"button",href:"docs/intro"},"Documentation")),n.createElement("div",{className:(0,o.Z)("col col--1")}),n.createElement("div",{className:(0,o.Z)("col col--6")},n.createElement("div",{className:"text--right"},n.createElement("img",{src:(0,c.Z)("img/banner.svg")}))))))};function u(){var e=(0,i.Z)().siteConfig;return n.createElement(r.Z,{title:e.tagline,description:"Meteor is an easy-to-use, plugin-driven metadata collection framework to extract data from different sources and sink to any data catalog or store."},n.createElement(d,null),n.createElement("main",null,n.createElement(s.Z,{className:"textSection wrapper",background:"light"},n.createElement("h1",null,"Built for ease"),n.createElement("p",null,"Meteor is a plugin driven agent for collecting metadata. Meteor has plugins to source metadata from a variety of data stores, services and message queues. It also has sink plugins to send metadata to variety of third party APIs and catalog services."),n.createElement(l.Z,{layout:"threeColumn",contents:[{title:"Zero dependency",content:n.createElement("div",null,"Meteor is written in Go and compiles into a single binary with no external dependencies, and requires a very minimal memory footprint.")},{title:"Coverage",content:n.createElement("div",null,"With 50+ plugins and many more coming soon to extract and sink metadata, it is easy to start collecting metadata from various sources.")},{title:"Extensible",content:n.createElement("div",null,"With the ease of plugin development you can build your own plugin to fit with your needs. It allows new sources, processors and sinks to be easily added.")},{title:"CLI",content:n.createElement("div",null,"Meteor comes with a CLI which allows you to interact with agent effectively. You can list all plugins, start and stop agent, and more.")},{title:"Proven",content:n.createElement("div",null,"Battle tested at large scale across multiple companies. Largest deployment collect metadata from thousands of data sources.")},{title:"Runtime",content:n.createElement("div",null,"Meteor can run from your local machine, cloud server machine or containers with minium efforts required for deployment.")}]})),n.createElement(s.Z,{className:"textSection wrapper",background:"dark"},n.createElement("h1",null,"Framework"),n.createElement("p",null,"Meteor agent uses recipes as a set of instructions which are configured by user. Recipes contains configurations about the source from which the metadata will be fetched, information about metadata processors and the destination to where the metadata will be sent."),n.createElement(l.Z,{layout:"threeColumn",contents:[{title:"Extraction",content:n.createElement("div",null,"Extraction is the process of extracting data from a source and transforming it into a format that can be consumed by the agent. Extractors are the set of plugins that are source of our metadata and include databases, dashboards, users, etc.")},{title:"Processing",content:n.createElement("div",null,"Processing is the process of transforming the extracted data into a format that can be consumed by the agent. Processors are the set of plugins that perform the enrichment or data processing for the metadata after extraction..")},{title:"Sink",content:n.createElement("div",null,"Sink is the process of sending the processed data to a single or multiple destinations as defined in recipes. Sinks are the set of plugins that act as the destination of our metadata after extraction and processing is done by agent.")}]})),n.createElement(s.Z,{className:"textSection wrapper",background:"light"},n.createElement("h1",null,"Ecosystem"),n.createElement("p",null,"Meteor\u2019s plugin system allows new plugins to be easily added. With 50+ plugins and many more coming soon to extract and sink metadata, it is easy to start collecting metadata from various sources and sink to any data catalog or store."),n.createElement("div",{className:"row"},n.createElement("div",{className:"col col--4"},n.createElement(l.Z,{contents:[{title:"Extractors",content:n.createElement("div",null,"Meteor supports source plugins to extract metadata from a variety of datastores services, and message queues, including BigQuery, InfluxDB, Kafka, Metabase, and many others.")},{title:"Processing",content:n.createElement("div",null,"Meteor has in-built processors inlcuding enrichment and others. It is easy to add your own processors as well using custom plugins.")},{title:"Sink",content:n.createElement("div",null,"Meteor supports sink plugins to send metadata to a variety of third party APIs and catalog services, including Columbus, HTTP, BigQuery, Kafka, and many others.")}]})),n.createElement("div",{className:"col col--8"},n.createElement("img",{src:(0,c.Z)("assets/overview_4.svg")}))))))}},4184:function(e,t){var a;!function(){"use strict";var n={}.hasOwnProperty;function r(){for(var e=[],t=0;t<arguments.length;t++){var a=arguments[t];if(a){var o=typeof a;if("string"===o||"number"===o)e.push(a);else if(Array.isArray(a)){if(a.length){var i=r.apply(null,a);i&&e.push(i)}}else if("object"===o)if(a.toString===Object.prototype.toString)for(var s in a)n.call(a,s)&&a[s]&&e.push(s);else e.push(a.toString())}}return e.join(" ")}e.exports?(r.default=r,e.exports=r):void 0===(a=function(){return r}.apply(t,[]))||(e.exports=a)}()}}]);