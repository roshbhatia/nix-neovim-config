" sysinit-nvim-core remote plugin stub
if has('nvim-0.5')
  " MANIFEST-BEGIN
  call remote#host#RegisterPlugin('sysinit-nvim-core', '0', [
  \ {'type': 'autocmd', 'name': 'VimEnter', 'sync': 1, 'opts': {'pattern': '*'}},
  \ ])
  " MANIFEST-END
endif
