export default {
  datafromutil: "datafromutil",
  testcolor: ["rgb(129, 212, 250)","rgb(239, 154, 154)","rgb(128, 203, 196)","rgb(129, 199, 132)","rgb(156, 204, 101)","rgb(212, 225, 87)","rgb(255, 213, 79)","rgb(255, 167, 38)","rgb(255, 110, 64)","rgb(161, 136, 127)","rgb(248, 187, 208)","rgb(128, 203, 196)","rgb(209, 196, 233)"],
  initial_style: (chart) => {
    console.log("图开始自适应")
    let chartParent = chart.parentNode.parentNode;
    return chartParent.clientWidth - 30 + "px";
  },
  set_color : () => {
    let testcolor = ["rgb(129, 212, 250)","rgb(239, 154, 154)","rgb(128, 203, 196)","rgb(129, 199, 132)","rgb(156, 204, 101)","rgb(212, 225, 87)","rgb(255, 213, 79)","rgb(255, 167, 38)","rgb(255, 110, 64)","rgb(161, 136, 127)","rgb(248, 187, 208)","rgb(128, 203, 196)","rgb(209, 196, 233)"]
    let num = Math.floor(Math.random() * testcolor.length);
    return testcolor[num]
  },
  sortNumber : (x,y) => {
    return x["value"]-y["value"];
  },
  get_detail: (name, detail_list) => {
    if (detail_list) {
      let str = "";
      detail_list.forEach(el => {
        if (el.name === name) {
          str += "NAME: " + name + "<br/>";
          for (let i in el) {
            if(i === "name") continue;
            str += i.toUpperCase() + ": ";
            str += (el[i] ? el[i] : " ");
            str += "<br/>";
          }
        }
      });
      if(!str) str = "NAME: " + name;
      return str;
    } else {
      return "NAME: " + name;
    }
  }
}