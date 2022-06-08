import "./style.scss";

const els = [
  document.querySelector("#app"),
  document.querySelector(".box-1"),
  document.querySelector(".box-2"),
];

document.addEventListener("mousemove", (e) => {
  els.forEach((layer) => {
    const s = layer.getAttribute("data-speed");
    const x = (window.innerWidth - e.pageX * s) / 100;
    const y = (window.innerHeight - e.pageY * s) / 100;
    layer.style.transform = `translate(${x}px, ${y}px)`;
  });
});