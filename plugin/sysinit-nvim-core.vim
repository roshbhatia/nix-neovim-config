" sysinit-nvim-core remote plugin stub
if has('nvim-0.5')
  " Plugin host factory and registration
  let s:host = 'sysinit-nvim-core'
  let s:bin = stdpath('config') . '/bin/' . s:host
  function! s:StartHost(host_info) abort
    return jobstart([s:bin], {'rpc': v:true})
  endfunction
  call remote#host#Register(s:host, '*', function('s:StartHost'))

  " MANIFEST-BEGIN
  call remote#host#RegisterPlugin(s:host, '0', [
  ])
  " MANIFEST-END
endif
