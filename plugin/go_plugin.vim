if exists('*remote#host#Register')
  if executable('go-plugin')
    call remote#host#Register('go-plugin', 'go-plugin')
  endif
endif