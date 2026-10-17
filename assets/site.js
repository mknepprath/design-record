(function () {
  var input = document.getElementById("filter");
  if (!input) return;
  var rows = document.querySelectorAll(".listing tbody tr");
  var none = document.getElementById("no-match");
  input.addEventListener("input", function () {
    var q = input.value.trim().toLowerCase();
    var shown = 0;
    rows.forEach(function (r) {
      var hit = !q || r.textContent.toLowerCase().indexOf(q) !== -1;
      r.hidden = !hit;
      if (hit) shown++;
    });
    document.querySelectorAll(".listing").forEach(function (s) {
      var any = s.querySelector("tbody tr:not([hidden])");
      s.hidden = !!q && !any;
    });
    none.hidden = shown > 0;
  });
})();
