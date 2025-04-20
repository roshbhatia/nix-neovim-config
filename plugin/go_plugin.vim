if exists('*remote#host#Register')
  " Prefer the locally built binary in nvim config bin directory
  let s:bin = stdpath('config') . '/bin/go-plugin'
  if filereadable(s:bin)
    call remote#host#Register('go-plugin', s:bin)
  elseif executable('go-plugin')
    call remote#host#Register('go-plugin', 'go-plugin')
  endif
endif