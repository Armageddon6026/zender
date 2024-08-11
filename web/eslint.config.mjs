import globals from 'globals'
import pluginJs from '@eslint/js'
import tseslint from 'typescript-eslint'
import pluginVue from 'eslint-plugin-vue'
import tsParser from '@typescript-eslint/parser'
import unocss from '@unocss/eslint-config/flat'

export default [
  {
    files: ['src/*', 'src/**/*.ts', 'src/**/*.vue', 'src/**/**/*.vue'],
    languageOptions: {
      globals: {
        ...globals.node,
        RequestModel: 'readonly',
        ResponseModel: 'readonly',
        EventModel: 'readonly',
        Dialog: 'readonly'
      },
      parserOptions: {
        ecmaFeatures: { modules: true },
        ecmaVersion: 'latest',
        parser: tsParser
      }
    }
  },
  pluginJs.configs.recommended,
  ...tseslint.configs.recommended,
  ...pluginVue.configs['flat/essential'],
  unocss,
  {
    rules: {
      'vue/multi-word-component-names': ['off'],
      '@typescript-eslint/no-explicit-any': ['off'],
      'no-unsafe-optional-chaining': ['warn'],
      'vue/valid-v-for': ['off']
    }
  }
]
