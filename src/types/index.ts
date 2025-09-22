export interface ConversionResult {
  cssModulesPath: string;
  componentPath: string;
  cssContent: string;
  updatedComponent: string;
  changeReport?: ChangeReport;
}

export interface ChangeReport {
  originalFile: string;
  modifiedFile: string;
  cssModulesFile: string;
  changes: Change[];
  summary: ChangeSummary;
}

export interface Change {
  type: 'class-replacement' | 'import-added';
  line: number;
  original: string;
  modified: string;
  cssClass: string;
  semanticName: string;
}

export interface ChangeSummary {
  totalNodes: number;
  classesConverted: number;
  importsAdded: number;
  cssRulesGenerated: number;
}

export interface CliArgs {
  input: string;
  output: string;
  verbose?: boolean;
  diff?: boolean;
  htmlDiff?: boolean;
  theme?: string;
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

export interface ThemeConfig {
  colors?: Record<string, any>;
  fontSize?: Record<string, any>;
  spacing?: Record<string, any>;
  borderRadius?: Record<string, any>;
  fontWeight?: Record<string, any>;
  [key: string]: any;
}

export interface ThemeLoadResult {
  theme: ThemeConfig;
  path: string;
} 