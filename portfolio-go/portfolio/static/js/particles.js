(function () {
  function initCanvas(cv) {
    var ctx = cv.getContext("2d");
    var W, H, pts = [];
    var N = 55, LINK = 130;

    function resize() {
      var p = cv.parentElement;
      W = cv.width  = p.offsetWidth  || window.innerWidth;
      H = cv.height = p.offsetHeight || 420;
    }

    function dot(scatter) {
      return {
        x:  Math.random() * W,
        y:  scatter ? Math.random() * H : H + 10,
        vx: (Math.random() - 0.5) * 0.4,
        vy: -(Math.random() * 0.5 + 0.15),
        r:  Math.random() * 2.5 + 1,
        a:  Math.random() * 0.5 + 0.2
      };
    }

    function frame() {
      ctx.clearRect(0, 0, W, H);
      for (var i = 0; i < pts.length; i++) {
        var p = pts[i];
        p.x += p.vx; p.y += p.vy;
        if (p.y < -10 || p.x < -10 || p.x > W + 10) { pts[i] = dot(false); continue; }
        for (var j = i + 1; j < pts.length; j++) {
          var q = pts[j];
          var dx = p.x - q.x, dy = p.y - q.y;
          var d = Math.sqrt(dx * dx + dy * dy);
          if (d < LINK) {
            ctx.globalAlpha = (1 - d / LINK) * 0.2;
            ctx.strokeStyle = "rgb(42,99,246)";
            ctx.lineWidth = 0.8;
            ctx.beginPath(); ctx.moveTo(p.x, p.y); ctx.lineTo(q.x, q.y); ctx.stroke();
          }
        }
        ctx.globalAlpha = p.a;
        ctx.fillStyle = "rgb(42,99,246)";
        ctx.beginPath(); ctx.arc(p.x, p.y, p.r, 0, 6.28); ctx.fill();
      }
      ctx.globalAlpha = 1;
      requestAnimationFrame(frame);
    }

    resize();
    for (var i = 0; i < N; i++) pts.push(dot(true));
    window.addEventListener("resize", resize);
    frame();
  }

  function init() {
    document.querySelectorAll("canvas[data-particles]").forEach(initCanvas);
  }

  if (document.readyState === "complete") { init(); }
  else { window.addEventListener("load", init); }
})();
