To detect the starting postion and velocity of the stone, we need also three hailstones. So we have $s$ and $h_1, h_2, h_3$, each containing position $x,y,z$ and velocity $dx,dy,dz$

For collision to happen we need time $t$ so that

$x_s + t*dx_s = x_{h1} + t*dx_{h1}$

$y_s + t*dy_s = y_{h1} + t*dy_{h1}$

$x_s + t*dx_s = x_{h1} + t*dx_{h1}$

Solving first two equasion for $t$

$t = (x_{h1} - x_s) / (dx_s - dx_{h1})$

$t = (y_{h1} - y_s) / (dy_s - dy_{h1})$

Which means that

$(x_{h1} - x_s) / (dx_s - dx_{h1}) 
= (y_{h1} - y_s) / (dy_s - dy_{h1}) =>$

$(x_{h1} - x_s) (dy_s - dy_{h1})  
= (y_{h1} - y_s) (dx_s - dx_{h1}) =>$

$x_{h1}dy_s - x_{h1}dy_{h1} - x_sdy_s + x_sdy_{h1}  
= y_{h1}dx_s - y_{h1}dx_{h1} - y_sdx_s + y_sdx_{h1} =>$

$x_{h1}dy_s - x_{h1}dy_{h1} - x_sdy_s + x_sdy_{h1}  
= y_{h1}dx_s - y_{h1}dx_{h1} - y_sdx_s + y_sdx_{h1} =>$

$ y_sdx_s - x_sdy_s
= y_{h1}dx_s - y_{h1}dx_{h1}  + y_sdx_{h1} - x_{h1}dy_s + x_{h1}dy_{h1} - x_sdy_{h1}$

Expression $y_sdx_s - x_sdy_s$ is constant for each hailstone, so we can take our second hailstone and get the same expression, and they must be equal for $h_1$ and $h_2$.

$y_{h1}dx_s - y_{h1}dx_{h1}  + y_sdx_{h1} - x_{h1}dy_s + x_{h1}dy_{h1} - x_sdy_{h1}
=y_{h2}dx_s - y_{h2}dx_{h2}  + y_sdx_{h2} - x_{h2}dy_s + x_{h2}dy_{h2} - x_sdy_{h2} =>$

$ (dy_{h2}-dy_{h1})x_s +  (y_{h1}-y_{h2})dx_s + 
(dx_{h1} - dx_{h2})y_s + (x_{h2} - x_{h1})dy_s
= x_{h2}dy_{h2} - y_{h2}dx_{h2} - x_{h1}dy_{h1} + y_{h1}dx_{h1}$

Now everything in the right side is known and left side looks very much like equasion for the system of linear equations. We have six unknowns, so we need 6 formulas, we have used $x$ and $y$ from $h_1$ and $h_2$. We can get two more equations using $x,z$ and $y,z$. And three more from $x,y,x$ combinations using $h1$ and $h3$. Final set of formulas needed looks like this:


$ (dy_{h2}-dy_{h1})x_s +  (y_{h1}-y_{h2})dx_s + 
(dx_{h1} - dx_{h2})y_s + (x_{h2} - x_{h1})dy_s
= x_{h2}dy_{h2} - y_{h2}dx_{h2} - x_{h1}dy_{h1} + y_{h1}dx_{h1}$

$ (dz_{h2}-dz_{h1})x_s +  (z_{h1}-z_{h2})dx_s + 
(dx_{h1} - dx_{h2})z_s + (x_{h2} - x_{h1})dz_s
= x_{h2}dz_{h2} - z_{h2}dx_{h2} - x_{h1}dz_{h1} + z_{h1}dx_{h1}$

$ (dz_{h2}-dz_{h1})y_s +  (z_{h1}-z_{h2})dy_s + 
(dy_{h1} - dy_{h2})z_s + (y_{h2} - y_{h1})dz_s
= y_{h2}dz_{h2} - z_{h2}dy_{h2} - y_{h1}dz_{h1} + z_{h1}dy_{h1}$

$ (dy_{h3}-dy_{h1})x_s +  (y_{h1}-y_{h3})dx_s + 
(dx_{h1} - dx_{h3})y_s + (x_{h3} - x_{h1})dy_s
= x_{h3}dy_{h3} - y_{h3}dx_{h3} - x_{h1}dy_{h1} + y_{h1}dx_{h1}$

$ (dz_{h3}-dz_{h1})x_s +  (z_{h1}-z_{h3})dx_s + 
(dx_{h1} - dx_{h3})z_s + (x_{h3} - x_{h1})dz_s
= x_{h3}dz_{h3} - z_{h3}dx_{h3} - x_{h1}dz_{h1} + z_{h1}dx_{h1}$

$ (dz_{h3}-dz_{h1})y_s +  (z_{h1}-z_{h3})dy_s + 
(dy_{h1} - dy_{h3})z_s + (y_{h3} - y_{h1})dz_s
= y_{h3}dz_{h3} - z_{h3}dy_{h3} - y_{h1}dz_{h1} + z_{h1}dy_{h1}$

Any combination of the three hailstones should work, so at least once we are swimming in overabundance of the data.
