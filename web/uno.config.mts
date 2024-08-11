import { defineConfig, presetAttributify, presetUno, presetTypography, presetIcons } from 'unocss'

export default defineConfig({
  shortcuts: {
    tds: 'b-solid border-0 border-gray pl',
    'zender-form': 'border-solid border-2px border-main-normal rd-10px'
  },
  theme: {
    colors: {
      main: {
        normal: 'var(--main-color)',
        dark: 'var(--main-color-dark)'
      },
      accept: { normal: 'var(--main-accept-color)', dark: 'var(--main-accept-color-dark)' },
      deny: {
        normal: 'var(--main-deny-color)',
        dark: 'var(--main-deny-color-dark)'
      }
    }
  },
  presets: [presetUno(), presetAttributify(), presetTypography(), presetIcons()]
})
