export interface ConversionResult {
  cssModulesPath: string;
  componentPath: string;
  cssContent: string;
  updatedComponent: string;
}

export interface CliArgs {
  input: string;
  output: string;
  verbose?: boolean;
}

export interface JSXNode {
  type: 'component' | 'html';
  tagName: string;
  className: string;
  line: number;
  semanticName: string;
}

export interface GeneratedCSS {
  selector: string;
  rules: string;
  modifiers?: CSSModifier[];
}

export interface CSSModifier {
  pseudo: string;
  rules: string;
}

export interface UnoGeneratorResult {
  css: string;
  matched: Set<string>;
} 