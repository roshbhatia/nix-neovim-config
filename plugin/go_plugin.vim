" Register the Go-based sysinit plugin as a remote host
" Remove the exists guard so autoload will define remote#host#Register
let s:bin = stdpath('config') . '/bin/sysinit-nvim-core'
if filereadable(s:bin)
  echom 'SysInit: registering remote host "sysinit-nvim-core" from ' . s:bin
  call remote#host#Register('sysinit-nvim-core', s:bin)
elseif executable('sysinit-nvim-core')
  echom 'SysInit: registering remote host "sysinit-nvim-core" from PATH'
  call remote#host#Register('sysinit-nvim-core', 'sysinit-nvim-core')
else
  echom 'SysInit: sysinit-nvim-core binary not found'
endif