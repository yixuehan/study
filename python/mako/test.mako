${"this is some text" | u}
<ul>
% for a in ("one", "two", "three"):
    <li>Item ${loop.index}: ${a}</li>
% endfor
</ul>
