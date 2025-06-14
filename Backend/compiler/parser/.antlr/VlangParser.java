// Generated from /home/brandon/Escritorio/OLC2_Proyecto1_202300813/Backend/compiler/parser/Vlang.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class VlangParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, TIPO=8, CASTEOS=9, 
		ATOI=10, PARSEFLOAT=11, TYPEOF=12, LEN=13, CAP=14, APPEND=15, IF=16, ELSE=17, 
		FOR=18, SWITCH=19, INDEXOF=20, JOIN=21, BREAK=22, CONTINUE=23, RETURN=24, 
		BOOLEANO=25, ENTERO=26, DECIMAL=27, CADENA=28, CARACTER=29, TIPO_ENTERO=30, 
		TIPO_DECIMAL=31, TIPO_CADENA=32, TIPO_BOOLEANO=33, TIPO_CHAR=34, PRINT=35, 
		ID=36, INC=37, DEC=38, SUMAIMPLICITA=39, RESTOIMPLICITO=40, PLUS=41, MINUS=42, 
		MUL=43, DIV=44, MOD=45, NOT=46, OR=47, AND=48, EQ=49, NEQ=50, LE=51, GE=52, 
		LT=53, GT=54, ASSIGN=55, LPAREN=56, RPAREN=57, LBRACK=58, RBRACK=59, LBRACE=60, 
		RBRACE=61, SEMICOLON=62, COLON=63, DOT=64, COMMA=65, WS=66, LINE_COMMENT=67, 
		BLOCK_COMMENT=68;
	public static final int
		RULE_programa = 0, RULE_funcMain = 1, RULE_funcDcl = 2, RULE_block = 3, 
		RULE_declaraciones = 4, RULE_varDcl = 5, RULE_sliceTipo = 6, RULE_sliceInit = 7, 
		RULE_stmt = 8, RULE_sentencias_control = 9, RULE_sentencias_transferencia = 10, 
		RULE_ifDcl = 11, RULE_elseIfDcl = 12, RULE_elseCondicional = 13, RULE_forDcl = 14, 
		RULE_asignacion = 15, RULE_switchDcl = 16, RULE_caseBlock = 17, RULE_defaultBlock = 18, 
		RULE_llamadaFuncion = 19, RULE_funcCall = 20, RULE_parametrosFormales = 21, 
		RULE_parametro = 22, RULE_parametrosReales = 23, RULE_structDcl = 24, 
		RULE_atributosStruct = 25, RULE_atributoStruct = 26, RULE_listaAsignaciones = 27, 
		RULE_asignacionStruct = 28, RULE_whileDcl = 29, RULE_expresion = 30, RULE_parametros = 31, 
		RULE_valores = 32, RULE_valor = 33, RULE_listaExpresiones = 34, RULE_incredecre = 35;
	private static String[] makeRuleNames() {
		return new String[] {
			"programa", "funcMain", "funcDcl", "block", "declaraciones", "varDcl", 
			"sliceTipo", "sliceInit", "stmt", "sentencias_control", "sentencias_transferencia", 
			"ifDcl", "elseIfDcl", "elseCondicional", "forDcl", "asignacion", "switchDcl", 
			"caseBlock", "defaultBlock", "llamadaFuncion", "funcCall", "parametrosFormales", 
			"parametro", "parametrosReales", "structDcl", "atributosStruct", "atributoStruct", 
			"listaAsignaciones", "asignacionStruct", "whileDcl", "expresion", "parametros", 
			"valores", "valor", "listaExpresiones", "incredecre"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'fn'", "'main'", "'mut'", "'case'", "'default'", "'struct'", "'while'", 
			null, null, "'Atoi'", "'parseFloat'", "'typeOf'", "'len'", "'cap'", "'append'", 
			"'if'", "'else'", "'for'", "'switch'", "'indexOf'", "'join'", "'break'", 
			"'continue'", "'return'", null, null, null, null, null, "'int'", "'float64'", 
			"'string'", "'bool'", "'rune'", "'print'", null, "'++'", "'--'", "'+='", 
			"'-='", "'+'", "'-'", "'*'", "'/'", "'%'", "'!'", "'||'", "'&&'", "'=='", 
			"'!='", "'<='", "'>='", "'<'", "'>'", "'='", "'('", "')'", "'['", "']'", 
			"'{'", "'}'", "';'", "':'", "'.'", "','"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, "TIPO", "CASTEOS", "ATOI", 
			"PARSEFLOAT", "TYPEOF", "LEN", "CAP", "APPEND", "IF", "ELSE", "FOR", 
			"SWITCH", "INDEXOF", "JOIN", "BREAK", "CONTINUE", "RETURN", "BOOLEANO", 
			"ENTERO", "DECIMAL", "CADENA", "CARACTER", "TIPO_ENTERO", "TIPO_DECIMAL", 
			"TIPO_CADENA", "TIPO_BOOLEANO", "TIPO_CHAR", "PRINT", "ID", "INC", "DEC", 
			"SUMAIMPLICITA", "RESTOIMPLICITO", "PLUS", "MINUS", "MUL", "DIV", "MOD", 
			"NOT", "OR", "AND", "EQ", "NEQ", "LE", "GE", "LT", "GT", "ASSIGN", "LPAREN", 
			"RPAREN", "LBRACK", "RBRACK", "LBRACE", "RBRACE", "SEMICOLON", "COLON", 
			"DOT", "COMMA", "WS", "LINE_COMMENT", "BLOCK_COMMENT"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "Vlang.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public VlangParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ProgramaContext extends ParserRuleContext {
		public TerminalNode EOF() { return getToken(VlangParser.EOF, 0); }
		public List<DeclaracionesContext> declaraciones() {
			return getRuleContexts(DeclaracionesContext.class);
		}
		public DeclaracionesContext declaraciones(int i) {
			return getRuleContext(DeclaracionesContext.class,i);
		}
		public ProgramaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_programa; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterPrograma(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitPrograma(this);
		}
	}

	public final ProgramaContext programa() throws RecognitionException {
		ProgramaContext _localctx = new ProgramaContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_programa);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(75);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 360362841133134026L) != 0)) {
				{
				{
				setState(72);
				declaraciones();
				}
				}
				setState(77);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(78);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class FuncMainContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(VlangParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(VlangParser.RPAREN, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public FuncMainContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funcMain; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterFuncMain(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitFuncMain(this);
		}
	}

	public final FuncMainContext funcMain() throws RecognitionException {
		FuncMainContext _localctx = new FuncMainContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_funcMain);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(80);
			match(T__0);
			setState(81);
			match(T__1);
			setState(82);
			match(LPAREN);
			setState(83);
			match(RPAREN);
			setState(84);
			block();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class FuncDclContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public TerminalNode LPAREN() { return getToken(VlangParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(VlangParser.RPAREN, 0); }
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public ParametrosFormalesContext parametrosFormales() {
			return getRuleContext(ParametrosFormalesContext.class,0);
		}
		public TerminalNode TIPO() { return getToken(VlangParser.TIPO, 0); }
		public FuncDclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funcDcl; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterFuncDcl(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitFuncDcl(this);
		}
	}

	public final FuncDclContext funcDcl() throws RecognitionException {
		FuncDclContext _localctx = new FuncDclContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_funcDcl);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(86);
			match(T__0);
			setState(87);
			match(ID);
			setState(88);
			match(LPAREN);
			setState(90);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ID) {
				{
				setState(89);
				parametrosFormales();
				}
			}

			setState(92);
			match(RPAREN);
			setState(94);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==TIPO) {
				{
				setState(93);
				match(TIPO);
				}
			}

			setState(96);
			block();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class BlockContext extends ParserRuleContext {
		public TerminalNode LBRACE() { return getToken(VlangParser.LBRACE, 0); }
		public TerminalNode RBRACE() { return getToken(VlangParser.RBRACE, 0); }
		public List<DeclaracionesContext> declaraciones() {
			return getRuleContexts(DeclaracionesContext.class);
		}
		public DeclaracionesContext declaraciones(int i) {
			return getRuleContext(DeclaracionesContext.class,i);
		}
		public BlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_block; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterBlock(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitBlock(this);
		}
	}

	public final BlockContext block() throws RecognitionException {
		BlockContext _localctx = new BlockContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_block);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(98);
			match(LBRACE);
			setState(102);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 360362841133134026L) != 0)) {
				{
				{
				setState(99);
				declaraciones();
				}
				}
				setState(104);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(105);
			match(RBRACE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class DeclaracionesContext extends ParserRuleContext {
		public VarDclContext varDcl() {
			return getRuleContext(VarDclContext.class,0);
		}
		public StmtContext stmt() {
			return getRuleContext(StmtContext.class,0);
		}
		public FuncDclContext funcDcl() {
			return getRuleContext(FuncDclContext.class,0);
		}
		public FuncMainContext funcMain() {
			return getRuleContext(FuncMainContext.class,0);
		}
		public StructDclContext structDcl() {
			return getRuleContext(StructDclContext.class,0);
		}
		public DeclaracionesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaraciones; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterDeclaraciones(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitDeclaraciones(this);
		}
	}

	public final DeclaracionesContext declaraciones() throws RecognitionException {
		DeclaracionesContext _localctx = new DeclaracionesContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_declaraciones);
		try {
			setState(112);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(107);
				varDcl();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(108);
				stmt();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(109);
				funcDcl();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(110);
				funcMain();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(111);
				structDcl();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class VarDclContext extends ParserRuleContext {
		public VarDclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_varDcl; }
	 
		public VarDclContext() { }
		public void copyFrom(VarDclContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StructDirectInitDeclarationContext extends VarDclContext {
		public List<TerminalNode> ID() { return getTokens(VlangParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(VlangParser.ID, i);
		}
		public TerminalNode ASSIGN() { return getToken(VlangParser.ASSIGN, 0); }
		public TerminalNode LBRACE() { return getToken(VlangParser.LBRACE, 0); }
		public ListaAsignacionesContext listaAsignaciones() {
			return getRuleContext(ListaAsignacionesContext.class,0);
		}
		public TerminalNode RBRACE() { return getToken(VlangParser.RBRACE, 0); }
		public StructDirectInitDeclarationContext(VarDclContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterStructDirectInitDeclaration(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitStructDirectInitDeclaration(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SliceEmptyDeclarationContext extends VarDclContext {
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public SliceTipoContext sliceTipo() {
			return getRuleContext(SliceTipoContext.class,0);
		}
		public SliceEmptyDeclarationContext(VarDclContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterSliceEmptyDeclaration(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitSliceEmptyDeclaration(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SliceAssignmentContext extends VarDclContext {
		public List<TerminalNode> ID() { return getTokens(VlangParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(VlangParser.ID, i);
		}
		public TerminalNode ASSIGN() { return getToken(VlangParser.ASSIGN, 0); }
		public SliceAssignmentContext(VarDclContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterSliceAssignment(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitSliceAssignment(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class VariableDeclarationImmutableContext extends VarDclContext {
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public TerminalNode ASSIGN() { return getToken(VlangParser.ASSIGN, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public VariableDeclarationImmutableContext(VarDclContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterVariableDeclarationImmutable(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitVariableDeclarationImmutable(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class VariableCastDeclarationContext extends VarDclContext {
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public TerminalNode ASSIGN() { return getToken(VlangParser.ASSIGN, 0); }
		public TerminalNode CASTEOS() { return getToken(VlangParser.CASTEOS, 0); }
		public TerminalNode LPAREN() { return getToken(VlangParser.LPAREN, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(VlangParser.RPAREN, 0); }
		public VariableCastDeclarationContext(VarDclContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterVariableCastDeclaration(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitVariableCastDeclaration(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SliceInitDeclarationContext extends VarDclContext {
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public TerminalNode ASSIGN() { return getToken(VlangParser.ASSIGN, 0); }
		public SliceTipoContext sliceTipo() {
			return getRuleContext(SliceTipoContext.class,0);
		}
		public SliceInitContext sliceInit() {
			return getRuleContext(SliceInitContext.class,0);
		}
		public SliceInitDeclarationContext(VarDclContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterSliceInitDeclaration(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitSliceInitDeclaration(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class VariableDeclarationContext extends VarDclContext {
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public TerminalNode TIPO() { return getToken(VlangParser.TIPO, 0); }
		public TerminalNode ASSIGN() { return getToken(VlangParser.ASSIGN, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public VariableDeclarationContext(VarDclContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterVariableDeclaration(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitVariableDeclaration(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SliceAssignmentIndexContext extends VarDclContext {
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public TerminalNode LBRACK() { return getToken(VlangParser.LBRACK, 0); }
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public TerminalNode RBRACK() { return getToken(VlangParser.RBRACK, 0); }
		public TerminalNode ASSIGN() { return getToken(VlangParser.ASSIGN, 0); }
		public SliceAssignmentIndexContext(VarDclContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterSliceAssignmentIndex(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitSliceAssignmentIndex(this);
		}
	}

	public final VarDclContext varDcl() throws RecognitionException {
		VarDclContext _localctx = new VarDclContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_varDcl);
		int _la;
		try {
			setState(160);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				_localctx = new VariableDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(114);
				match(T__2);
				setState(115);
				match(ID);
				setState(117);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==TIPO) {
					{
					setState(116);
					match(TIPO);
					}
				}

				setState(121);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==ASSIGN) {
					{
					setState(119);
					match(ASSIGN);
					setState(120);
					expresion(0);
					}
				}

				}
				break;
			case 2:
				_localctx = new SliceEmptyDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(123);
				match(T__2);
				setState(124);
				match(ID);
				setState(125);
				sliceTipo();
				}
				break;
			case 3:
				_localctx = new StructDirectInitDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(126);
				match(ID);
				setState(127);
				match(ASSIGN);
				setState(128);
				match(ID);
				setState(129);
				match(LBRACE);
				setState(130);
				listaAsignaciones();
				setState(131);
				match(RBRACE);
				}
				break;
			case 4:
				_localctx = new SliceInitDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(133);
				match(ID);
				setState(134);
				match(ASSIGN);
				setState(135);
				sliceTipo();
				setState(136);
				sliceInit();
				}
				break;
			case 5:
				_localctx = new SliceAssignmentContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(138);
				match(ID);
				setState(139);
				match(ASSIGN);
				setState(140);
				match(ID);
				}
				break;
			case 6:
				_localctx = new VariableDeclarationImmutableContext(_localctx);
				enterOuterAlt(_localctx, 6);
				{
				setState(141);
				match(ID);
				setState(144);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==ASSIGN) {
					{
					setState(142);
					match(ASSIGN);
					setState(143);
					expresion(0);
					}
				}

				}
				break;
			case 7:
				_localctx = new VariableCastDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 7);
				{
				setState(146);
				match(ID);
				setState(147);
				match(ASSIGN);
				setState(148);
				match(CASTEOS);
				setState(149);
				match(LPAREN);
				setState(150);
				expresion(0);
				setState(151);
				match(RPAREN);
				}
				break;
			case 8:
				_localctx = new SliceAssignmentIndexContext(_localctx);
				enterOuterAlt(_localctx, 8);
				{
				setState(153);
				match(ID);
				setState(154);
				match(LBRACK);
				setState(155);
				expresion(0);
				setState(156);
				match(RBRACK);
				setState(157);
				match(ASSIGN);
				setState(158);
				expresion(0);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SliceTipoContext extends ParserRuleContext {
		public TerminalNode LBRACK() { return getToken(VlangParser.LBRACK, 0); }
		public TerminalNode RBRACK() { return getToken(VlangParser.RBRACK, 0); }
		public TerminalNode TIPO() { return getToken(VlangParser.TIPO, 0); }
		public SliceTipoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sliceTipo; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterSliceTipo(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitSliceTipo(this);
		}
	}

	public final SliceTipoContext sliceTipo() throws RecognitionException {
		SliceTipoContext _localctx = new SliceTipoContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_sliceTipo);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(162);
			match(LBRACK);
			setState(163);
			match(RBRACK);
			setState(164);
			match(TIPO);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SliceInitContext extends ParserRuleContext {
		public TerminalNode LBRACE() { return getToken(VlangParser.LBRACE, 0); }
		public TerminalNode RBRACE() { return getToken(VlangParser.RBRACE, 0); }
		public ListaExpresionesContext listaExpresiones() {
			return getRuleContext(ListaExpresionesContext.class,0);
		}
		public SliceInitContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sliceInit; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterSliceInit(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitSliceInit(this);
		}
	}

	public final SliceInitContext sliceInit() throws RecognitionException {
		SliceInitContext _localctx = new SliceInitContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_sliceInit);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(166);
			match(LBRACE);
			setState(168);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 360362806743183360L) != 0)) {
				{
				setState(167);
				listaExpresiones();
				}
			}

			setState(170);
			match(RBRACE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StmtContext extends ParserRuleContext {
		public StmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_stmt; }
	 
		public StmtContext() { }
		public void copyFrom(StmtContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class PrintStatementContext extends StmtContext {
		public TerminalNode PRINT() { return getToken(VlangParser.PRINT, 0); }
		public TerminalNode LPAREN() { return getToken(VlangParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(VlangParser.RPAREN, 0); }
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(VlangParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(VlangParser.COMMA, i);
		}
		public PrintStatementContext(StmtContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterPrintStatement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitPrintStatement(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ControlStatementContext extends StmtContext {
		public Sentencias_controlContext sentencias_control() {
			return getRuleContext(Sentencias_controlContext.class,0);
		}
		public ControlStatementContext(StmtContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterControlStatement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitControlStatement(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExpresionStatementContext extends StmtContext {
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public ExpresionStatementContext(StmtContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterExpresionStatement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitExpresionStatement(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class TransfersentenceContext extends StmtContext {
		public Sentencias_transferenciaContext sentencias_transferencia() {
			return getRuleContext(Sentencias_transferenciaContext.class,0);
		}
		public TransfersentenceContext(StmtContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterTransfersentence(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitTransfersentence(this);
		}
	}

	public final StmtContext stmt() throws RecognitionException {
		StmtContext _localctx = new StmtContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_stmt);
		int _la;
		try {
			setState(188);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case PRINT:
				_localctx = new PrintStatementContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(172);
				match(PRINT);
				setState(173);
				match(LPAREN);
				setState(182);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 360362806743183360L) != 0)) {
					{
					setState(174);
					expresion(0);
					setState(179);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==COMMA) {
						{
						{
						setState(175);
						match(COMMA);
						setState(176);
						expresion(0);
						}
						}
						setState(181);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					}
				}

				setState(184);
				match(RPAREN);
				}
				break;
			case TYPEOF:
			case LEN:
			case APPEND:
			case INDEXOF:
			case JOIN:
			case BOOLEANO:
			case ENTERO:
			case DECIMAL:
			case CADENA:
			case CARACTER:
			case ID:
			case MINUS:
			case NOT:
			case LPAREN:
			case LBRACK:
				_localctx = new ExpresionStatementContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(185);
				expresion(0);
				}
				break;
			case T__6:
			case IF:
			case FOR:
			case SWITCH:
				_localctx = new ControlStatementContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(186);
				sentencias_control();
				}
				break;
			case BREAK:
			case CONTINUE:
			case RETURN:
				_localctx = new TransfersentenceContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(187);
				sentencias_transferencia();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Sentencias_controlContext extends ParserRuleContext {
		public Sentencias_controlContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sentencias_control; }
	 
		public Sentencias_controlContext() { }
		public void copyFrom(Sentencias_controlContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class While_contextContext extends Sentencias_controlContext {
		public WhileDclContext whileDcl() {
			return getRuleContext(WhileDclContext.class,0);
		}
		public While_contextContext(Sentencias_controlContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterWhile_context(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitWhile_context(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Switch_contextContext extends Sentencias_controlContext {
		public SwitchDclContext switchDcl() {
			return getRuleContext(SwitchDclContext.class,0);
		}
		public Switch_contextContext(Sentencias_controlContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterSwitch_context(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitSwitch_context(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class If_contextContext extends Sentencias_controlContext {
		public IfDclContext ifDcl() {
			return getRuleContext(IfDclContext.class,0);
		}
		public If_contextContext(Sentencias_controlContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterIf_context(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitIf_context(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class For_contextContext extends Sentencias_controlContext {
		public ForDclContext forDcl() {
			return getRuleContext(ForDclContext.class,0);
		}
		public For_contextContext(Sentencias_controlContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterFor_context(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitFor_context(this);
		}
	}

	public final Sentencias_controlContext sentencias_control() throws RecognitionException {
		Sentencias_controlContext _localctx = new Sentencias_controlContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_sentencias_control);
		try {
			setState(194);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case IF:
				_localctx = new If_contextContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(190);
				ifDcl();
				}
				break;
			case FOR:
				_localctx = new For_contextContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(191);
				forDcl();
				}
				break;
			case SWITCH:
				_localctx = new Switch_contextContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(192);
				switchDcl();
				}
				break;
			case T__6:
				_localctx = new While_contextContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(193);
				whileDcl();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Sentencias_transferenciaContext extends ParserRuleContext {
		public Sentencias_transferenciaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sentencias_transferencia; }
	 
		public Sentencias_transferenciaContext() { }
		public void copyFrom(Sentencias_transferenciaContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class BreakStatementContext extends Sentencias_transferenciaContext {
		public TerminalNode BREAK() { return getToken(VlangParser.BREAK, 0); }
		public BreakStatementContext(Sentencias_transferenciaContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterBreakStatement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitBreakStatement(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ContinueStatementContext extends Sentencias_transferenciaContext {
		public TerminalNode CONTINUE() { return getToken(VlangParser.CONTINUE, 0); }
		public ContinueStatementContext(Sentencias_transferenciaContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterContinueStatement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitContinueStatement(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ReturnStatementContext extends Sentencias_transferenciaContext {
		public TerminalNode RETURN() { return getToken(VlangParser.RETURN, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public ReturnStatementContext(Sentencias_transferenciaContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterReturnStatement(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitReturnStatement(this);
		}
	}

	public final Sentencias_transferenciaContext sentencias_transferencia() throws RecognitionException {
		Sentencias_transferenciaContext _localctx = new Sentencias_transferenciaContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_sentencias_transferencia);
		try {
			setState(202);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case BREAK:
				_localctx = new BreakStatementContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(196);
				match(BREAK);
				}
				break;
			case CONTINUE:
				_localctx = new ContinueStatementContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(197);
				match(CONTINUE);
				}
				break;
			case RETURN:
				_localctx = new ReturnStatementContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(198);
				match(RETURN);
				setState(200);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,14,_ctx) ) {
				case 1:
					{
					setState(199);
					expresion(0);
					}
					break;
				}
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class IfDclContext extends ParserRuleContext {
		public TerminalNode IF() { return getToken(VlangParser.IF, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode LBRACE() { return getToken(VlangParser.LBRACE, 0); }
		public TerminalNode RBRACE() { return getToken(VlangParser.RBRACE, 0); }
		public List<DeclaracionesContext> declaraciones() {
			return getRuleContexts(DeclaracionesContext.class);
		}
		public DeclaracionesContext declaraciones(int i) {
			return getRuleContext(DeclaracionesContext.class,i);
		}
		public List<ElseIfDclContext> elseIfDcl() {
			return getRuleContexts(ElseIfDclContext.class);
		}
		public ElseIfDclContext elseIfDcl(int i) {
			return getRuleContext(ElseIfDclContext.class,i);
		}
		public ElseCondicionalContext elseCondicional() {
			return getRuleContext(ElseCondicionalContext.class,0);
		}
		public IfDclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ifDcl; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterIfDcl(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitIfDcl(this);
		}
	}

	public final IfDclContext ifDcl() throws RecognitionException {
		IfDclContext _localctx = new IfDclContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_ifDcl);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(204);
			match(IF);
			setState(205);
			expresion(0);
			setState(206);
			match(LBRACE);
			setState(210);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 360362841133134026L) != 0)) {
				{
				{
				setState(207);
				declaraciones();
				}
				}
				setState(212);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(213);
			match(RBRACE);
			setState(217);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,17,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(214);
					elseIfDcl();
					}
					} 
				}
				setState(219);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,17,_ctx);
			}
			setState(221);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ELSE) {
				{
				setState(220);
				elseCondicional();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ElseIfDclContext extends ParserRuleContext {
		public TerminalNode ELSE() { return getToken(VlangParser.ELSE, 0); }
		public TerminalNode IF() { return getToken(VlangParser.IF, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode LBRACE() { return getToken(VlangParser.LBRACE, 0); }
		public TerminalNode RBRACE() { return getToken(VlangParser.RBRACE, 0); }
		public List<DeclaracionesContext> declaraciones() {
			return getRuleContexts(DeclaracionesContext.class);
		}
		public DeclaracionesContext declaraciones(int i) {
			return getRuleContext(DeclaracionesContext.class,i);
		}
		public ElseIfDclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_elseIfDcl; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterElseIfDcl(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitElseIfDcl(this);
		}
	}

	public final ElseIfDclContext elseIfDcl() throws RecognitionException {
		ElseIfDclContext _localctx = new ElseIfDclContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_elseIfDcl);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(223);
			match(ELSE);
			setState(224);
			match(IF);
			setState(225);
			expresion(0);
			setState(226);
			match(LBRACE);
			setState(230);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 360362841133134026L) != 0)) {
				{
				{
				setState(227);
				declaraciones();
				}
				}
				setState(232);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(233);
			match(RBRACE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ElseCondicionalContext extends ParserRuleContext {
		public TerminalNode ELSE() { return getToken(VlangParser.ELSE, 0); }
		public TerminalNode LBRACE() { return getToken(VlangParser.LBRACE, 0); }
		public TerminalNode RBRACE() { return getToken(VlangParser.RBRACE, 0); }
		public List<DeclaracionesContext> declaraciones() {
			return getRuleContexts(DeclaracionesContext.class);
		}
		public DeclaracionesContext declaraciones(int i) {
			return getRuleContext(DeclaracionesContext.class,i);
		}
		public ElseCondicionalContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_elseCondicional; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterElseCondicional(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitElseCondicional(this);
		}
	}

	public final ElseCondicionalContext elseCondicional() throws RecognitionException {
		ElseCondicionalContext _localctx = new ElseCondicionalContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_elseCondicional);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(235);
			match(ELSE);
			setState(236);
			match(LBRACE);
			setState(240);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 360362841133134026L) != 0)) {
				{
				{
				setState(237);
				declaraciones();
				}
				}
				setState(242);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(243);
			match(RBRACE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ForDclContext extends ParserRuleContext {
		public ForDclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_forDcl; }
	 
		public ForDclContext() { }
		public void copyFrom(ForDclContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ForCondicionUnicaContext extends ForDclContext {
		public TerminalNode FOR() { return getToken(VlangParser.FOR, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public ForCondicionUnicaContext(ForDclContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterForCondicionUnica(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitForCondicionUnica(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ForClasicoContext extends ForDclContext {
		public TerminalNode FOR() { return getToken(VlangParser.FOR, 0); }
		public AsignacionContext asignacion() {
			return getRuleContext(AsignacionContext.class,0);
		}
		public List<TerminalNode> SEMICOLON() { return getTokens(VlangParser.SEMICOLON); }
		public TerminalNode SEMICOLON(int i) {
			return getToken(VlangParser.SEMICOLON, i);
		}
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public StmtContext stmt() {
			return getRuleContext(StmtContext.class,0);
		}
		public ForClasicoContext(ForDclContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterForClasico(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitForClasico(this);
		}
	}

	public final ForDclContext forDcl() throws RecognitionException {
		ForDclContext _localctx = new ForDclContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_forDcl);
		int _la;
		try {
			setState(259);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,22,_ctx) ) {
			case 1:
				_localctx = new ForClasicoContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(245);
				match(FOR);
				setState(246);
				asignacion();
				setState(247);
				match(SEMICOLON);
				setState(248);
				expresion(0);
				setState(249);
				match(SEMICOLON);
				setState(251);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 360362841133133952L) != 0)) {
					{
					setState(250);
					stmt();
					}
				}

				setState(253);
				block();
				}
				break;
			case 2:
				_localctx = new ForCondicionUnicaContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(255);
				match(FOR);
				setState(256);
				expresion(0);
				setState(257);
				block();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class AsignacionContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public TerminalNode ASSIGN() { return getToken(VlangParser.ASSIGN, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public AsignacionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_asignacion; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterAsignacion(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitAsignacion(this);
		}
	}

	public final AsignacionContext asignacion() throws RecognitionException {
		AsignacionContext _localctx = new AsignacionContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_asignacion);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(261);
			match(ID);
			setState(262);
			match(ASSIGN);
			setState(263);
			expresion(0);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class SwitchDclContext extends ParserRuleContext {
		public TerminalNode SWITCH() { return getToken(VlangParser.SWITCH, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode LBRACE() { return getToken(VlangParser.LBRACE, 0); }
		public TerminalNode RBRACE() { return getToken(VlangParser.RBRACE, 0); }
		public List<CaseBlockContext> caseBlock() {
			return getRuleContexts(CaseBlockContext.class);
		}
		public CaseBlockContext caseBlock(int i) {
			return getRuleContext(CaseBlockContext.class,i);
		}
		public DefaultBlockContext defaultBlock() {
			return getRuleContext(DefaultBlockContext.class,0);
		}
		public SwitchDclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_switchDcl; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterSwitchDcl(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitSwitchDcl(this);
		}
	}

	public final SwitchDclContext switchDcl() throws RecognitionException {
		SwitchDclContext _localctx = new SwitchDclContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_switchDcl);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(265);
			match(SWITCH);
			setState(266);
			expresion(0);
			setState(267);
			match(LBRACE);
			setState(271);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__3) {
				{
				{
				setState(268);
				caseBlock();
				}
				}
				setState(273);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(275);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__4) {
				{
				setState(274);
				defaultBlock();
				}
			}

			setState(277);
			match(RBRACE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class CaseBlockContext extends ParserRuleContext {
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode COLON() { return getToken(VlangParser.COLON, 0); }
		public List<DeclaracionesContext> declaraciones() {
			return getRuleContexts(DeclaracionesContext.class);
		}
		public DeclaracionesContext declaraciones(int i) {
			return getRuleContext(DeclaracionesContext.class,i);
		}
		public CaseBlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_caseBlock; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterCaseBlock(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitCaseBlock(this);
		}
	}

	public final CaseBlockContext caseBlock() throws RecognitionException {
		CaseBlockContext _localctx = new CaseBlockContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_caseBlock);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(279);
			match(T__3);
			setState(280);
			expresion(0);
			setState(281);
			match(COLON);
			setState(285);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 360362841133134026L) != 0)) {
				{
				{
				setState(282);
				declaraciones();
				}
				}
				setState(287);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class DefaultBlockContext extends ParserRuleContext {
		public TerminalNode COLON() { return getToken(VlangParser.COLON, 0); }
		public List<DeclaracionesContext> declaraciones() {
			return getRuleContexts(DeclaracionesContext.class);
		}
		public DeclaracionesContext declaraciones(int i) {
			return getRuleContext(DeclaracionesContext.class,i);
		}
		public DefaultBlockContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_defaultBlock; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterDefaultBlock(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitDefaultBlock(this);
		}
	}

	public final DefaultBlockContext defaultBlock() throws RecognitionException {
		DefaultBlockContext _localctx = new DefaultBlockContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_defaultBlock);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(288);
			match(T__4);
			setState(289);
			match(COLON);
			setState(293);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 360362841133134026L) != 0)) {
				{
				{
				setState(290);
				declaraciones();
				}
				}
				setState(295);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class LlamadaFuncionContext extends ParserRuleContext {
		public TerminalNode INDEXOF() { return getToken(VlangParser.INDEXOF, 0); }
		public TerminalNode LPAREN() { return getToken(VlangParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(VlangParser.RPAREN, 0); }
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(VlangParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(VlangParser.COMMA, i);
		}
		public TerminalNode JOIN() { return getToken(VlangParser.JOIN, 0); }
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public TerminalNode LEN() { return getToken(VlangParser.LEN, 0); }
		public TerminalNode APPEND() { return getToken(VlangParser.APPEND, 0); }
		public LlamadaFuncionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_llamadaFuncion; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterLlamadaFuncion(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitLlamadaFuncion(this);
		}
	}

	public final LlamadaFuncionContext llamadaFuncion() throws RecognitionException {
		LlamadaFuncionContext _localctx = new LlamadaFuncionContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_llamadaFuncion);
		int _la;
		try {
			setState(361);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INDEXOF:
				enterOuterAlt(_localctx, 1);
				{
				setState(296);
				match(INDEXOF);
				setState(297);
				match(LPAREN);
				setState(306);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 360362806743183360L) != 0)) {
					{
					setState(298);
					expresion(0);
					setState(303);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==COMMA) {
						{
						{
						setState(299);
						match(COMMA);
						setState(300);
						expresion(0);
						}
						}
						setState(305);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					}
				}

				setState(308);
				match(RPAREN);
				}
				break;
			case JOIN:
				enterOuterAlt(_localctx, 2);
				{
				setState(309);
				match(JOIN);
				setState(310);
				match(LPAREN);
				setState(319);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 360362806743183360L) != 0)) {
					{
					setState(311);
					expresion(0);
					setState(316);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==COMMA) {
						{
						{
						setState(312);
						match(COMMA);
						setState(313);
						expresion(0);
						}
						}
						setState(318);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					}
				}

				setState(321);
				match(RPAREN);
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 3);
				{
				setState(322);
				match(ID);
				setState(323);
				match(LPAREN);
				setState(332);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 360362806743183360L) != 0)) {
					{
					setState(324);
					expresion(0);
					setState(329);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==COMMA) {
						{
						{
						setState(325);
						match(COMMA);
						setState(326);
						expresion(0);
						}
						}
						setState(331);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					}
				}

				setState(334);
				match(RPAREN);
				}
				break;
			case LEN:
				enterOuterAlt(_localctx, 4);
				{
				setState(335);
				match(LEN);
				setState(336);
				match(LPAREN);
				setState(345);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 360362806743183360L) != 0)) {
					{
					setState(337);
					expresion(0);
					setState(342);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==COMMA) {
						{
						{
						setState(338);
						match(COMMA);
						setState(339);
						expresion(0);
						}
						}
						setState(344);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					}
				}

				setState(347);
				match(RPAREN);
				}
				break;
			case APPEND:
				enterOuterAlt(_localctx, 5);
				{
				setState(348);
				match(APPEND);
				setState(349);
				match(LPAREN);
				setState(358);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 360362806743183360L) != 0)) {
					{
					setState(350);
					expresion(0);
					setState(355);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==COMMA) {
						{
						{
						setState(351);
						match(COMMA);
						setState(352);
						expresion(0);
						}
						}
						setState(357);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					}
				}

				setState(360);
				match(RPAREN);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class FuncCallContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public TerminalNode LPAREN() { return getToken(VlangParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(VlangParser.RPAREN, 0); }
		public ParametrosRealesContext parametrosReales() {
			return getRuleContext(ParametrosRealesContext.class,0);
		}
		public FuncCallContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funcCall; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterFuncCall(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitFuncCall(this);
		}
	}

	public final FuncCallContext funcCall() throws RecognitionException {
		FuncCallContext _localctx = new FuncCallContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_funcCall);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(363);
			match(ID);
			setState(364);
			match(LPAREN);
			setState(366);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 360362806743183360L) != 0)) {
				{
				setState(365);
				parametrosReales();
				}
			}

			setState(368);
			match(RPAREN);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ParametrosFormalesContext extends ParserRuleContext {
		public List<ParametroContext> parametro() {
			return getRuleContexts(ParametroContext.class);
		}
		public ParametroContext parametro(int i) {
			return getRuleContext(ParametroContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(VlangParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(VlangParser.COMMA, i);
		}
		public ParametrosFormalesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parametrosFormales; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterParametrosFormales(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitParametrosFormales(this);
		}
	}

	public final ParametrosFormalesContext parametrosFormales() throws RecognitionException {
		ParametrosFormalesContext _localctx = new ParametrosFormalesContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_parametrosFormales);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(370);
			parametro();
			setState(375);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(371);
				match(COMMA);
				setState(372);
				parametro();
				}
				}
				setState(377);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ParametroContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public TerminalNode TIPO() { return getToken(VlangParser.TIPO, 0); }
		public ParametroContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parametro; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterParametro(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitParametro(this);
		}
	}

	public final ParametroContext parametro() throws RecognitionException {
		ParametroContext _localctx = new ParametroContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_parametro);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(378);
			match(ID);
			setState(379);
			match(TIPO);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ParametrosRealesContext extends ParserRuleContext {
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(VlangParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(VlangParser.COMMA, i);
		}
		public ParametrosRealesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parametrosReales; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterParametrosReales(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitParametrosReales(this);
		}
	}

	public final ParametrosRealesContext parametrosReales() throws RecognitionException {
		ParametrosRealesContext _localctx = new ParametrosRealesContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_parametrosReales);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(381);
			expresion(0);
			setState(386);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(382);
				match(COMMA);
				setState(383);
				expresion(0);
				}
				}
				setState(388);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StructDclContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public TerminalNode LBRACE() { return getToken(VlangParser.LBRACE, 0); }
		public AtributosStructContext atributosStruct() {
			return getRuleContext(AtributosStructContext.class,0);
		}
		public TerminalNode RBRACE() { return getToken(VlangParser.RBRACE, 0); }
		public StructDclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_structDcl; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterStructDcl(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitStructDcl(this);
		}
	}

	public final StructDclContext structDcl() throws RecognitionException {
		StructDclContext _localctx = new StructDclContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_structDcl);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(389);
			match(T__5);
			setState(390);
			match(ID);
			setState(391);
			match(LBRACE);
			setState(392);
			atributosStruct();
			setState(393);
			match(RBRACE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class AtributosStructContext extends ParserRuleContext {
		public List<AtributoStructContext> atributoStruct() {
			return getRuleContexts(AtributoStructContext.class);
		}
		public AtributoStructContext atributoStruct(int i) {
			return getRuleContext(AtributoStructContext.class,i);
		}
		public AtributosStructContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_atributosStruct; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterAtributosStruct(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitAtributosStruct(this);
		}
	}

	public final AtributosStructContext atributosStruct() throws RecognitionException {
		AtributosStructContext _localctx = new AtributosStructContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_atributosStruct);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(396); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(395);
				atributoStruct();
				}
				}
				setState(398); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==ID );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class AtributoStructContext extends ParserRuleContext {
		public AtributoStructContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_atributoStruct; }
	 
		public AtributoStructContext() { }
		public void copyFrom(AtributoStructContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AtributoPrimitivoContext extends AtributoStructContext {
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public TerminalNode TIPO() { return getToken(VlangParser.TIPO, 0); }
		public AtributoPrimitivoContext(AtributoStructContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterAtributoPrimitivo(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitAtributoPrimitivo(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AtributoStructAnidadoContext extends AtributoStructContext {
		public List<TerminalNode> ID() { return getTokens(VlangParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(VlangParser.ID, i);
		}
		public AtributoStructAnidadoContext(AtributoStructContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterAtributoStructAnidado(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitAtributoStructAnidado(this);
		}
	}

	public final AtributoStructContext atributoStruct() throws RecognitionException {
		AtributoStructContext _localctx = new AtributoStructContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_atributoStruct);
		try {
			setState(404);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,42,_ctx) ) {
			case 1:
				_localctx = new AtributoPrimitivoContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(400);
				match(ID);
				setState(401);
				match(TIPO);
				}
				break;
			case 2:
				_localctx = new AtributoStructAnidadoContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(402);
				match(ID);
				setState(403);
				match(ID);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ListaAsignacionesContext extends ParserRuleContext {
		public List<AsignacionStructContext> asignacionStruct() {
			return getRuleContexts(AsignacionStructContext.class);
		}
		public AsignacionStructContext asignacionStruct(int i) {
			return getRuleContext(AsignacionStructContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(VlangParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(VlangParser.COMMA, i);
		}
		public ListaAsignacionesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listaAsignaciones; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterListaAsignaciones(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitListaAsignaciones(this);
		}
	}

	public final ListaAsignacionesContext listaAsignaciones() throws RecognitionException {
		ListaAsignacionesContext _localctx = new ListaAsignacionesContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_listaAsignaciones);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(406);
			asignacionStruct();
			setState(411);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(407);
				match(COMMA);
				setState(408);
				asignacionStruct();
				}
				}
				setState(413);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class AsignacionStructContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public TerminalNode COLON() { return getToken(VlangParser.COLON, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public AsignacionStructContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_asignacionStruct; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterAsignacionStruct(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitAsignacionStruct(this);
		}
	}

	public final AsignacionStructContext asignacionStruct() throws RecognitionException {
		AsignacionStructContext _localctx = new AsignacionStructContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_asignacionStruct);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(414);
			match(ID);
			setState(415);
			match(COLON);
			setState(416);
			expresion(0);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class WhileDclContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(VlangParser.LPAREN, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(VlangParser.RPAREN, 0); }
		public TerminalNode LBRACK() { return getToken(VlangParser.LBRACK, 0); }
		public TerminalNode RBRACK() { return getToken(VlangParser.RBRACK, 0); }
		public List<DeclaracionesContext> declaraciones() {
			return getRuleContexts(DeclaracionesContext.class);
		}
		public DeclaracionesContext declaraciones(int i) {
			return getRuleContext(DeclaracionesContext.class,i);
		}
		public WhileDclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_whileDcl; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterWhileDcl(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitWhileDcl(this);
		}
	}

	public final WhileDclContext whileDcl() throws RecognitionException {
		WhileDclContext _localctx = new WhileDclContext(_ctx, getState());
		enterRule(_localctx, 58, RULE_whileDcl);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(418);
			match(T__6);
			setState(419);
			match(LPAREN);
			setState(420);
			expresion(0);
			setState(421);
			match(RPAREN);
			setState(422);
			match(LBRACK);
			setState(426);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 360362841133134026L) != 0)) {
				{
				{
				setState(423);
				declaraciones();
				}
				}
				setState(428);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(429);
			match(RBRACK);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ExpresionContext extends ParserRuleContext {
		public ExpresionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expresion; }
	 
		public ExpresionContext() { }
		public void copyFrom(ExpresionContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class MultdivmodContext extends ExpresionContext {
		public Token op;
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public TerminalNode MUL() { return getToken(VlangParser.MUL, 0); }
		public TerminalNode DIV() { return getToken(VlangParser.DIV, 0); }
		public TerminalNode MOD() { return getToken(VlangParser.MOD, 0); }
		public MultdivmodContext(ExpresionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterMultdivmod(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitMultdivmod(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Casteo_paratipoContext extends ExpresionContext {
		public TerminalNode TYPEOF() { return getToken(VlangParser.TYPEOF, 0); }
		public TerminalNode LPAREN() { return getToken(VlangParser.LPAREN, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(VlangParser.RPAREN, 0); }
		public Casteo_paratipoContext(ExpresionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterCasteo_paratipo(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitCasteo_paratipo(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IncredecrContext extends ExpresionContext {
		public IncredecreContext incredecre() {
			return getRuleContext(IncredecreContext.class,0);
		}
		public IncredecrContext(ExpresionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterIncredecr(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitIncredecr(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class OPERADORESLOGICOSContext extends ExpresionContext {
		public Token op;
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public TerminalNode AND() { return getToken(VlangParser.AND, 0); }
		public TerminalNode OR() { return getToken(VlangParser.OR, 0); }
		public OPERADORESLOGICOSContext(ExpresionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterOPERADORESLOGICOS(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitOPERADORESLOGICOS(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StructInstanceCreationContext extends ExpresionContext {
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public TerminalNode LBRACE() { return getToken(VlangParser.LBRACE, 0); }
		public ListaAsignacionesContext listaAsignaciones() {
			return getRuleContext(ListaAsignacionesContext.class,0);
		}
		public TerminalNode RBRACE() { return getToken(VlangParser.RBRACE, 0); }
		public StructInstanceCreationContext(ExpresionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterStructInstanceCreation(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitStructInstanceCreation(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ValorexprContext extends ExpresionContext {
		public ValorContext valor() {
			return getRuleContext(ValorContext.class,0);
		}
		public ValorexprContext(ExpresionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterValorexpr(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitValorexpr(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IgualdadContext extends ExpresionContext {
		public Token op;
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public TerminalNode EQ() { return getToken(VlangParser.EQ, 0); }
		public TerminalNode NEQ() { return getToken(VlangParser.NEQ, 0); }
		public IgualdadContext(ExpresionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterIgualdad(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitIgualdad(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class LlamadaFuncionExprContext extends ExpresionContext {
		public LlamadaFuncionContext llamadaFuncion() {
			return getRuleContext(LlamadaFuncionContext.class,0);
		}
		public LlamadaFuncionExprContext(ExpresionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterLlamadaFuncionExpr(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitLlamadaFuncionExpr(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExpdotexpContext extends ExpresionContext {
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public TerminalNode DOT() { return getToken(VlangParser.DOT, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public ExpdotexpContext(ExpresionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterExpdotexp(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitExpdotexp(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StructAttrAssignContext extends ExpresionContext {
		public List<TerminalNode> ID() { return getTokens(VlangParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(VlangParser.ID, i);
		}
		public TerminalNode DOT() { return getToken(VlangParser.DOT, 0); }
		public TerminalNode ASSIGN() { return getToken(VlangParser.ASSIGN, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public StructAttrAssignContext(ExpresionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterStructAttrAssign(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitStructAttrAssign(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class RelacionalesContext extends ExpresionContext {
		public Token op;
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public TerminalNode LT() { return getToken(VlangParser.LT, 0); }
		public TerminalNode LE() { return getToken(VlangParser.LE, 0); }
		public TerminalNode GE() { return getToken(VlangParser.GE, 0); }
		public TerminalNode GT() { return getToken(VlangParser.GT, 0); }
		public RelacionalesContext(ExpresionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterRelacionales(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitRelacionales(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Casteo_paratipo_sliceContext extends ExpresionContext {
		public TerminalNode TYPEOF() { return getToken(VlangParser.TYPEOF, 0); }
		public TerminalNode LPAREN() { return getToken(VlangParser.LPAREN, 0); }
		public TerminalNode LBRACK() { return getToken(VlangParser.LBRACK, 0); }
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public TerminalNode RBRACK() { return getToken(VlangParser.RBRACK, 0); }
		public TerminalNode RPAREN() { return getToken(VlangParser.RPAREN, 0); }
		public List<TerminalNode> COMMA() { return getTokens(VlangParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(VlangParser.COMMA, i);
		}
		public Casteo_paratipo_sliceContext(ExpresionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterCasteo_paratipo_slice(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitCasteo_paratipo_slice(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class CorchetesexpreContext extends ExpresionContext {
		public TerminalNode LBRACK() { return getToken(VlangParser.LBRACK, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode RBRACK() { return getToken(VlangParser.RBRACK, 0); }
		public CorchetesexpreContext(ExpresionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterCorchetesexpre(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitCorchetesexpre(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class UnarioContext extends ExpresionContext {
		public Token op;
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode NOT() { return getToken(VlangParser.NOT, 0); }
		public TerminalNode MINUS() { return getToken(VlangParser.MINUS, 0); }
		public UnarioContext(ExpresionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterUnario(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitUnario(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ParentesisexpreContext extends ExpresionContext {
		public TerminalNode LPAREN() { return getToken(VlangParser.LPAREN, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(VlangParser.RPAREN, 0); }
		public ParentesisexpreContext(ExpresionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterParentesisexpre(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitParentesisexpre(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IMCPLICITContext extends ExpresionContext {
		public Token op;
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode SUMAIMPLICITA() { return getToken(VlangParser.SUMAIMPLICITA, 0); }
		public TerminalNode RESTOIMPLICITO() { return getToken(VlangParser.RESTOIMPLICITO, 0); }
		public IMCPLICITContext(ExpresionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterIMCPLICIT(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitIMCPLICIT(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SumresContext extends ExpresionContext {
		public Token op;
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public TerminalNode PLUS() { return getToken(VlangParser.PLUS, 0); }
		public TerminalNode MINUS() { return getToken(VlangParser.MINUS, 0); }
		public SumresContext(ExpresionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterSumres(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitSumres(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class PARAPRINTSLICEContext extends ExpresionContext {
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public TerminalNode LBRACK() { return getToken(VlangParser.LBRACK, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode RBRACK() { return getToken(VlangParser.RBRACK, 0); }
		public PARAPRINTSLICEContext(ExpresionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterPARAPRINTSLICE(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitPARAPRINTSLICE(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AsignacionLUEGOContext extends ExpresionContext {
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public TerminalNode TIPO() { return getToken(VlangParser.TIPO, 0); }
		public TerminalNode ASSIGN() { return getToken(VlangParser.ASSIGN, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public AsignacionLUEGOContext(ExpresionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterAsignacionLUEGO(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitAsignacionLUEGO(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IdContext extends ExpresionContext {
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public IdContext(ExpresionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterId(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitId(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Expdotexp1Context extends ExpresionContext {
		public List<TerminalNode> ID() { return getTokens(VlangParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(VlangParser.ID, i);
		}
		public TerminalNode DOT() { return getToken(VlangParser.DOT, 0); }
		public Expdotexp1Context(ExpresionContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterExpdotexp1(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitExpdotexp1(this);
		}
	}

	public final ExpresionContext expresion() throws RecognitionException {
		return expresion(0);
	}

	private ExpresionContext expresion(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExpresionContext _localctx = new ExpresionContext(_ctx, _parentState);
		ExpresionContext _prevctx = _localctx;
		int _startState = 60;
		enterRecursionRule(_localctx, 60, RULE_expresion, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(493);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,46,_ctx) ) {
			case 1:
				{
				_localctx = new UnarioContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(432);
				((UnarioContext)_localctx).op = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==MINUS || _la==NOT) ) {
					((UnarioContext)_localctx).op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(433);
				expresion(19);
				}
				break;
			case 2:
				{
				_localctx = new ValorexprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(434);
				valor();
				}
				break;
			case 3:
				{
				_localctx = new ParentesisexpreContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(435);
				match(LPAREN);
				setState(436);
				expresion(0);
				setState(437);
				match(RPAREN);
				}
				break;
			case 4:
				{
				_localctx = new CorchetesexpreContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(439);
				match(LBRACK);
				setState(440);
				expresion(0);
				setState(441);
				match(RBRACK);
				}
				break;
			case 5:
				{
				_localctx = new PARAPRINTSLICEContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(443);
				match(ID);
				setState(444);
				match(LBRACK);
				setState(445);
				expresion(0);
				setState(446);
				match(RBRACK);
				}
				break;
			case 6:
				{
				_localctx = new LlamadaFuncionExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(448);
				llamadaFuncion();
				}
				break;
			case 7:
				{
				_localctx = new StructAttrAssignContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(449);
				match(ID);
				setState(450);
				match(DOT);
				setState(451);
				match(ID);
				setState(452);
				match(ASSIGN);
				setState(453);
				expresion(10);
				}
				break;
			case 8:
				{
				_localctx = new StructInstanceCreationContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(454);
				match(ID);
				setState(455);
				match(LBRACE);
				setState(456);
				listaAsignaciones();
				setState(457);
				match(RBRACE);
				}
				break;
			case 9:
				{
				_localctx = new IdContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(459);
				match(ID);
				}
				break;
			case 10:
				{
				_localctx = new IncredecrContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(460);
				incredecre();
				}
				break;
			case 11:
				{
				_localctx = new Expdotexp1Context(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(461);
				match(ID);
				setState(462);
				match(DOT);
				setState(463);
				match(ID);
				}
				break;
			case 12:
				{
				_localctx = new ExpdotexpContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(464);
				match(ID);
				setState(465);
				match(DOT);
				setState(466);
				expresion(5);
				}
				break;
			case 13:
				{
				_localctx = new AsignacionLUEGOContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(467);
				match(ID);
				setState(468);
				match(TIPO);
				setState(469);
				match(ASSIGN);
				setState(470);
				expresion(4);
				}
				break;
			case 14:
				{
				_localctx = new IMCPLICITContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(471);
				match(ID);
				setState(472);
				((IMCPLICITContext)_localctx).op = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==SUMAIMPLICITA || _la==RESTOIMPLICITO) ) {
					((IMCPLICITContext)_localctx).op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(473);
				expresion(3);
				}
				break;
			case 15:
				{
				_localctx = new Casteo_paratipo_sliceContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(474);
				match(TYPEOF);
				setState(475);
				match(LPAREN);
				setState(476);
				match(LBRACK);
				setState(477);
				expresion(0);
				setState(482);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==COMMA) {
					{
					{
					setState(478);
					match(COMMA);
					setState(479);
					expresion(0);
					}
					}
					setState(484);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(485);
				match(RBRACK);
				setState(486);
				match(RPAREN);
				}
				break;
			case 16:
				{
				_localctx = new Casteo_paratipoContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(488);
				match(TYPEOF);
				setState(489);
				match(LPAREN);
				setState(490);
				expresion(0);
				setState(491);
				match(RPAREN);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(512);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,48,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(510);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,47,_ctx) ) {
					case 1:
						{
						_localctx = new MultdivmodContext(new ExpresionContext(_parentctx, _parentState));
						_localctx = new MultdivmodContext(new ExpresionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(495);
						if (!(precpred(_ctx, 21))) throw new FailedPredicateException(this, "precpred(_ctx, 21)");
						setState(496);
						((MultdivmodContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 61572651155456L) != 0)) ) {
							((MultdivmodContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 61572651155456L) != 0)) ) {
							((MultdivmodContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(497);
						expresion(22);
						}
						break;
					case 2:
						{
						_localctx = new SumresContext(new ExpresionContext(_parentctx, _parentState));
						_localctx = new SumresContext(new ExpresionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(498);
						if (!(precpred(_ctx, 20))) throw new FailedPredicateException(this, "precpred(_ctx, 20)");
						setState(499);
						((SumresContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==PLUS || _la==MINUS) ) {
							((SumresContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						if ( !(_la==PLUS || _la==MINUS) ) {
							((SumresContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(500);
						expresion(21);
						}
						break;
					case 3:
						{
						_localctx = new RelacionalesContext(new ExpresionContext(_parentctx, _parentState));
						_localctx = new RelacionalesContext(new ExpresionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(501);
						if (!(precpred(_ctx, 18))) throw new FailedPredicateException(this, "precpred(_ctx, 18)");
						setState(502);
						((RelacionalesContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 33776997205278720L) != 0)) ) {
							((RelacionalesContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 33776997205278720L) != 0)) ) {
							((RelacionalesContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(503);
						expresion(19);
						}
						break;
					case 4:
						{
						_localctx = new IgualdadContext(new ExpresionContext(_parentctx, _parentState));
						_localctx = new IgualdadContext(new ExpresionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(504);
						if (!(precpred(_ctx, 17))) throw new FailedPredicateException(this, "precpred(_ctx, 17)");
						setState(505);
						((IgualdadContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==EQ || _la==NEQ) ) {
							((IgualdadContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						if ( !(_la==EQ || _la==NEQ) ) {
							((IgualdadContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(506);
						expresion(18);
						}
						break;
					case 5:
						{
						_localctx = new OPERADORESLOGICOSContext(new ExpresionContext(_parentctx, _parentState));
						_localctx = new OPERADORESLOGICOSContext(new ExpresionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(507);
						if (!(precpred(_ctx, 16))) throw new FailedPredicateException(this, "precpred(_ctx, 16)");
						setState(508);
						((OPERADORESLOGICOSContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==OR || _la==AND) ) {
							((OPERADORESLOGICOSContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						if ( !(_la==OR || _la==AND) ) {
							((OPERADORESLOGICOSContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(509);
						expresion(17);
						}
						break;
					}
					} 
				}
				setState(514);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,48,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ParametrosContext extends ParserRuleContext {
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(VlangParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(VlangParser.COMMA, i);
		}
		public ParametrosContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_parametros; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterParametros(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitParametros(this);
		}
	}

	public final ParametrosContext parametros() throws RecognitionException {
		ParametrosContext _localctx = new ParametrosContext(_ctx, getState());
		enterRule(_localctx, 62, RULE_parametros);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(515);
			expresion(0);
			setState(520);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(516);
				match(COMMA);
				setState(517);
				expresion(0);
				}
				}
				setState(522);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ValoresContext extends ParserRuleContext {
		public ValorContext valor() {
			return getRuleContext(ValorContext.class,0);
		}
		public ValoresContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_valores; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterValores(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitValores(this);
		}
	}

	public final ValoresContext valores() throws RecognitionException {
		ValoresContext _localctx = new ValoresContext(_ctx, getState());
		enterRule(_localctx, 64, RULE_valores);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(523);
			valor();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ValorContext extends ParserRuleContext {
		public ValorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_valor; }
	 
		public ValorContext() { }
		public void copyFrom(ValorContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ValorDecimalContext extends ValorContext {
		public TerminalNode DECIMAL() { return getToken(VlangParser.DECIMAL, 0); }
		public ValorDecimalContext(ValorContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterValorDecimal(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitValorDecimal(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ValorEnteroContext extends ValorContext {
		public TerminalNode ENTERO() { return getToken(VlangParser.ENTERO, 0); }
		public ValorEnteroContext(ValorContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterValorEntero(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitValorEntero(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ValorBooleanoContext extends ValorContext {
		public TerminalNode BOOLEANO() { return getToken(VlangParser.BOOLEANO, 0); }
		public ValorBooleanoContext(ValorContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterValorBooleano(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitValorBooleano(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ValorCaracterContext extends ValorContext {
		public TerminalNode CARACTER() { return getToken(VlangParser.CARACTER, 0); }
		public ValorCaracterContext(ValorContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterValorCaracter(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitValorCaracter(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ValorCadenaContext extends ValorContext {
		public TerminalNode CADENA() { return getToken(VlangParser.CADENA, 0); }
		public ValorCadenaContext(ValorContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterValorCadena(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitValorCadena(this);
		}
	}

	public final ValorContext valor() throws RecognitionException {
		ValorContext _localctx = new ValorContext(_ctx, getState());
		enterRule(_localctx, 66, RULE_valor);
		try {
			setState(530);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case ENTERO:
				_localctx = new ValorEnteroContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(525);
				match(ENTERO);
				}
				break;
			case DECIMAL:
				_localctx = new ValorDecimalContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(526);
				match(DECIMAL);
				}
				break;
			case CADENA:
				_localctx = new ValorCadenaContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(527);
				match(CADENA);
				}
				break;
			case BOOLEANO:
				_localctx = new ValorBooleanoContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(528);
				match(BOOLEANO);
				}
				break;
			case CARACTER:
				_localctx = new ValorCaracterContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(529);
				match(CARACTER);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ListaExpresionesContext extends ParserRuleContext {
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(VlangParser.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(VlangParser.COMMA, i);
		}
		public ListaExpresionesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listaExpresiones; }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterListaExpresiones(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitListaExpresiones(this);
		}
	}

	public final ListaExpresionesContext listaExpresiones() throws RecognitionException {
		ListaExpresionesContext _localctx = new ListaExpresionesContext(_ctx, getState());
		enterRule(_localctx, 68, RULE_listaExpresiones);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(532);
			expresion(0);
			setState(537);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(533);
				match(COMMA);
				setState(534);
				expresion(0);
				}
				}
				setState(539);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class IncredecreContext extends ParserRuleContext {
		public IncredecreContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_incredecre; }
	 
		public IncredecreContext() { }
		public void copyFrom(IncredecreContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IncrementoContext extends IncredecreContext {
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public TerminalNode INC() { return getToken(VlangParser.INC, 0); }
		public IncrementoContext(IncredecreContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterIncremento(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitIncremento(this);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DecrementoContext extends IncredecreContext {
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public TerminalNode DEC() { return getToken(VlangParser.DEC, 0); }
		public DecrementoContext(IncredecreContext ctx) { copyFrom(ctx); }
		@Override
		public void enterRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).enterDecremento(this);
		}
		@Override
		public void exitRule(ParseTreeListener listener) {
			if ( listener instanceof VlangListener ) ((VlangListener)listener).exitDecremento(this);
		}
	}

	public final IncredecreContext incredecre() throws RecognitionException {
		IncredecreContext _localctx = new IncredecreContext(_ctx, getState());
		enterRule(_localctx, 70, RULE_incredecre);
		try {
			setState(544);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,52,_ctx) ) {
			case 1:
				_localctx = new IncrementoContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(540);
				match(ID);
				setState(541);
				match(INC);
				}
				break;
			case 2:
				_localctx = new DecrementoContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(542);
				match(ID);
				setState(543);
				match(DEC);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 30:
			return expresion_sempred((ExpresionContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expresion_sempred(ExpresionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 21);
		case 1:
			return precpred(_ctx, 20);
		case 2:
			return precpred(_ctx, 18);
		case 3:
		case 3:
			return precpred(_ctx, 17);
		case 4:
			return precpred(_ctx, 16);
		case 4:
			return precpred(_ctx, 16);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001D\u0223\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002\u0012\u0007\u0012"+
		"\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0002\u0015\u0007\u0015"+
		"\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017\u0002\u0018\u0007\u0018"+
		"\u0002\u0019\u0007\u0019\u0002\u001a\u0007\u001a\u0002\u001b\u0007\u001b"+
		"\u0002\u001c\u0007\u001c\u0002\u001d\u0007\u001d\u0002\u001e\u0007\u001e"+
		"\u0002\u001f\u0007\u001f\u0002 \u0007 \u0002!\u0007!\u0002\"\u0007\"\u0002"+
		"#\u0007#\u0001\u0000\u0005\u0000J\b\u0000\n\u0000\f\u0000M\t\u0000\u0001"+
		"\u0000\u0001\u0000\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0003"+
		"\u0002[\b\u0002\u0001\u0002\u0001\u0002\u0003\u0002_\b\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0003\u0001\u0003\u0005\u0003e\b\u0003\n\u0003\f\u0003"+
		"h\t\u0003\u0001\u0003\u0001\u0003\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0003\u0004q\b\u0004\u0001\u0005\u0001\u0005"+
		"\u0001\u0005\u0003\u0005v\b\u0005\u0001\u0005\u0001\u0005\u0003\u0005"+
		"z\b\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0003\u0005\u0091\b\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0003\u0005\u00a1\b\u0005\u0001\u0006\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0001\u0007\u0001\u0007\u0003\u0007\u00a9\b\u0007"+
		"\u0001\u0007\u0001\u0007\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0005"+
		"\b\u00b2\b\b\n\b\f\b\u00b5\t\b\u0003\b\u00b7\b\b\u0001\b\u0001\b\u0001"+
		"\b\u0001\b\u0003\b\u00bd\b\b\u0001\t\u0001\t\u0001\t\u0001\t\u0003\t\u00c3"+
		"\b\t\u0001\n\u0001\n\u0001\n\u0001\n\u0003\n\u00c9\b\n\u0003\n\u00cb\b"+
		"\n\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0005\u000b\u00d1\b"+
		"\u000b\n\u000b\f\u000b\u00d4\t\u000b\u0001\u000b\u0001\u000b\u0005\u000b"+
		"\u00d8\b\u000b\n\u000b\f\u000b\u00db\t\u000b\u0001\u000b\u0003\u000b\u00de"+
		"\b\u000b\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0005\f\u00e5\b\f\n\f"+
		"\f\f\u00e8\t\f\u0001\f\u0001\f\u0001\r\u0001\r\u0001\r\u0005\r\u00ef\b"+
		"\r\n\r\f\r\u00f2\t\r\u0001\r\u0001\r\u0001\u000e\u0001\u000e\u0001\u000e"+
		"\u0001\u000e\u0001\u000e\u0001\u000e\u0003\u000e\u00fc\b\u000e\u0001\u000e"+
		"\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0003\u000e"+
		"\u0104\b\u000e\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u0010"+
		"\u0001\u0010\u0001\u0010\u0001\u0010\u0005\u0010\u010e\b\u0010\n\u0010"+
		"\f\u0010\u0111\t\u0010\u0001\u0010\u0003\u0010\u0114\b\u0010\u0001\u0010"+
		"\u0001\u0010\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0005\u0011"+
		"\u011c\b\u0011\n\u0011\f\u0011\u011f\t\u0011\u0001\u0012\u0001\u0012\u0001"+
		"\u0012\u0005\u0012\u0124\b\u0012\n\u0012\f\u0012\u0127\t\u0012\u0001\u0013"+
		"\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0005\u0013\u012e\b\u0013"+
		"\n\u0013\f\u0013\u0131\t\u0013\u0003\u0013\u0133\b\u0013\u0001\u0013\u0001"+
		"\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0005\u0013\u013b"+
		"\b\u0013\n\u0013\f\u0013\u013e\t\u0013\u0003\u0013\u0140\b\u0013\u0001"+
		"\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0005"+
		"\u0013\u0148\b\u0013\n\u0013\f\u0013\u014b\t\u0013\u0003\u0013\u014d\b"+
		"\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001"+
		"\u0013\u0005\u0013\u0155\b\u0013\n\u0013\f\u0013\u0158\t\u0013\u0003\u0013"+
		"\u015a\b\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013"+
		"\u0001\u0013\u0005\u0013\u0162\b\u0013\n\u0013\f\u0013\u0165\t\u0013\u0003"+
		"\u0013\u0167\b\u0013\u0001\u0013\u0003\u0013\u016a\b\u0013\u0001\u0014"+
		"\u0001\u0014\u0001\u0014\u0003\u0014\u016f\b\u0014\u0001\u0014\u0001\u0014"+
		"\u0001\u0015\u0001\u0015\u0001\u0015\u0005\u0015\u0176\b\u0015\n\u0015"+
		"\f\u0015\u0179\t\u0015\u0001\u0016\u0001\u0016\u0001\u0016\u0001\u0017"+
		"\u0001\u0017\u0001\u0017\u0005\u0017\u0181\b\u0017\n\u0017\f\u0017\u0184"+
		"\t\u0017\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001"+
		"\u0018\u0001\u0019\u0004\u0019\u018d\b\u0019\u000b\u0019\f\u0019\u018e"+
		"\u0001\u001a\u0001\u001a\u0001\u001a\u0001\u001a\u0003\u001a\u0195\b\u001a"+
		"\u0001\u001b\u0001\u001b\u0001\u001b\u0005\u001b\u019a\b\u001b\n\u001b"+
		"\f\u001b\u019d\t\u001b\u0001\u001c\u0001\u001c\u0001\u001c\u0001\u001c"+
		"\u0001\u001d\u0001\u001d\u0001\u001d\u0001\u001d\u0001\u001d\u0001\u001d"+
		"\u0005\u001d\u01a9\b\u001d\n\u001d\f\u001d\u01ac\t\u001d\u0001\u001d\u0001"+
		"\u001d\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001"+
		"\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001"+
		"\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001"+
		"\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001"+
		"\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001"+
		"\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001"+
		"\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001"+
		"\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001"+
		"\u001e\u0001\u001e\u0005\u001e\u01e1\b\u001e\n\u001e\f\u001e\u01e4\t\u001e"+
		"\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e"+
		"\u0001\u001e\u0001\u001e\u0003\u001e\u01ee\b\u001e\u0001\u001e\u0001\u001e"+
		"\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e"+
		"\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e\u0001\u001e"+
		"\u0001\u001e\u0005\u001e\u01ff\b\u001e\n\u001e\f\u001e\u0202\t\u001e\u0001"+
		"\u001f\u0001\u001f\u0001\u001f\u0005\u001f\u0207\b\u001f\n\u001f\f\u001f"+
		"\u020a\t\u001f\u0001 \u0001 \u0001!\u0001!\u0001!\u0001!\u0001!\u0003"+
		"!\u0213\b!\u0001\"\u0001\"\u0001\"\u0005\"\u0218\b\"\n\"\f\"\u021b\t\""+
		"\u0001#\u0001#\u0001#\u0001#\u0003#\u0221\b#\u0001#\u0000\u0001<$\u0000"+
		"\u0002\u0004\u0006\b\n\f\u000e\u0010\u0012\u0014\u0016\u0018\u001a\u001c"+
		"\u001e \"$&(*,.02468:<>@BDF\u0000\u0007\u0002\u0000**..\u0001\u0000\'"+
		"(\u0001\u0000+-\u0001\u0000)*\u0001\u000036\u0001\u000012\u0001\u0000"+
		"/0\u0258\u0000K\u0001\u0000\u0000\u0000\u0002P\u0001\u0000\u0000\u0000"+
		"\u0004V\u0001\u0000\u0000\u0000\u0006b\u0001\u0000\u0000\u0000\bp\u0001"+
		"\u0000\u0000\u0000\n\u00a0\u0001\u0000\u0000\u0000\f\u00a2\u0001\u0000"+
		"\u0000\u0000\u000e\u00a6\u0001\u0000\u0000\u0000\u0010\u00bc\u0001\u0000"+
		"\u0000\u0000\u0012\u00c2\u0001\u0000\u0000\u0000\u0014\u00ca\u0001\u0000"+
		"\u0000\u0000\u0016\u00cc\u0001\u0000\u0000\u0000\u0018\u00df\u0001\u0000"+
		"\u0000\u0000\u001a\u00eb\u0001\u0000\u0000\u0000\u001c\u0103\u0001\u0000"+
		"\u0000\u0000\u001e\u0105\u0001\u0000\u0000\u0000 \u0109\u0001\u0000\u0000"+
		"\u0000\"\u0117\u0001\u0000\u0000\u0000$\u0120\u0001\u0000\u0000\u0000"+
		"&\u0169\u0001\u0000\u0000\u0000(\u016b\u0001\u0000\u0000\u0000*\u0172"+
		"\u0001\u0000\u0000\u0000,\u017a\u0001\u0000\u0000\u0000.\u017d\u0001\u0000"+
		"\u0000\u00000\u0185\u0001\u0000\u0000\u00002\u018c\u0001\u0000\u0000\u0000"+
		"4\u0194\u0001\u0000\u0000\u00006\u0196\u0001\u0000\u0000\u00008\u019e"+
		"\u0001\u0000\u0000\u0000:\u01a2\u0001\u0000\u0000\u0000<\u01ed\u0001\u0000"+
		"\u0000\u0000>\u0203\u0001\u0000\u0000\u0000@\u020b\u0001\u0000\u0000\u0000"+
		"B\u0212\u0001\u0000\u0000\u0000D\u0214\u0001\u0000\u0000\u0000F\u0220"+
		"\u0001\u0000\u0000\u0000HJ\u0003\b\u0004\u0000IH\u0001\u0000\u0000\u0000"+
		"JM\u0001\u0000\u0000\u0000KI\u0001\u0000\u0000\u0000KL\u0001\u0000\u0000"+
		"\u0000LN\u0001\u0000\u0000\u0000MK\u0001\u0000\u0000\u0000NO\u0005\u0000"+
		"\u0000\u0001O\u0001\u0001\u0000\u0000\u0000PQ\u0005\u0001\u0000\u0000"+
		"QR\u0005\u0002\u0000\u0000RS\u00058\u0000\u0000ST\u00059\u0000\u0000T"+
		"U\u0003\u0006\u0003\u0000U\u0003\u0001\u0000\u0000\u0000VW\u0005\u0001"+
		"\u0000\u0000WX\u0005$\u0000\u0000XZ\u00058\u0000\u0000Y[\u0003*\u0015"+
		"\u0000ZY\u0001\u0000\u0000\u0000Z[\u0001\u0000\u0000\u0000[\\\u0001\u0000"+
		"\u0000\u0000\\^\u00059\u0000\u0000]_\u0005\b\u0000\u0000^]\u0001\u0000"+
		"\u0000\u0000^_\u0001\u0000\u0000\u0000_`\u0001\u0000\u0000\u0000`a\u0003"+
		"\u0006\u0003\u0000a\u0005\u0001\u0000\u0000\u0000bf\u0005<\u0000\u0000"+
		"ce\u0003\b\u0004\u0000dc\u0001\u0000\u0000\u0000eh\u0001\u0000\u0000\u0000"+
		"fd\u0001\u0000\u0000\u0000fg\u0001\u0000\u0000\u0000gi\u0001\u0000\u0000"+
		"\u0000hf\u0001\u0000\u0000\u0000ij\u0005=\u0000\u0000j\u0007\u0001\u0000"+
		"\u0000\u0000kq\u0003\n\u0005\u0000lq\u0003\u0010\b\u0000mq\u0003\u0004"+
		"\u0002\u0000nq\u0003\u0002\u0001\u0000oq\u00030\u0018\u0000pk\u0001\u0000"+
		"\u0000\u0000pl\u0001\u0000\u0000\u0000pm\u0001\u0000\u0000\u0000pn\u0001"+
		"\u0000\u0000\u0000po\u0001\u0000\u0000\u0000q\t\u0001\u0000\u0000\u0000"+
		"rs\u0005\u0003\u0000\u0000su\u0005$\u0000\u0000tv\u0005\b\u0000\u0000"+
		"ut\u0001\u0000\u0000\u0000uv\u0001\u0000\u0000\u0000vy\u0001\u0000\u0000"+
		"\u0000wx\u00057\u0000\u0000xz\u0003<\u001e\u0000yw\u0001\u0000\u0000\u0000"+
		"yz\u0001\u0000\u0000\u0000z\u00a1\u0001\u0000\u0000\u0000{|\u0005\u0003"+
		"\u0000\u0000|}\u0005$\u0000\u0000}\u00a1\u0003\f\u0006\u0000~\u007f\u0005"+
		"$\u0000\u0000\u007f\u0080\u00057\u0000\u0000\u0080\u0081\u0005$\u0000"+
		"\u0000\u0081\u0082\u0005<\u0000\u0000\u0082\u0083\u00036\u001b\u0000\u0083"+
		"\u0084\u0005=\u0000\u0000\u0084\u00a1\u0001\u0000\u0000\u0000\u0085\u0086"+
		"\u0005$\u0000\u0000\u0086\u0087\u00057\u0000\u0000\u0087\u0088\u0003\f"+
		"\u0006\u0000\u0088\u0089\u0003\u000e\u0007\u0000\u0089\u00a1\u0001\u0000"+
		"\u0000\u0000\u008a\u008b\u0005$\u0000\u0000\u008b\u008c\u00057\u0000\u0000"+
		"\u008c\u00a1\u0005$\u0000\u0000\u008d\u0090\u0005$\u0000\u0000\u008e\u008f"+
		"\u00057\u0000\u0000\u008f\u0091\u0003<\u001e\u0000\u0090\u008e\u0001\u0000"+
		"\u0000\u0000\u0090\u0091\u0001\u0000\u0000\u0000\u0091\u00a1\u0001\u0000"+
		"\u0000\u0000\u0092\u0093\u0005$\u0000\u0000\u0093\u0094\u00057\u0000\u0000"+
		"\u0094\u0095\u0005\t\u0000\u0000\u0095\u0096\u00058\u0000\u0000\u0096"+
		"\u0097\u0003<\u001e\u0000\u0097\u0098\u00059\u0000\u0000\u0098\u00a1\u0001"+
		"\u0000\u0000\u0000\u0099\u009a\u0005$\u0000\u0000\u009a\u009b\u0005:\u0000"+
		"\u0000\u009b\u009c\u0003<\u001e\u0000\u009c\u009d\u0005;\u0000\u0000\u009d"+
		"\u009e\u00057\u0000\u0000\u009e\u009f\u0003<\u001e\u0000\u009f\u00a1\u0001"+
		"\u0000\u0000\u0000\u00a0r\u0001\u0000\u0000\u0000\u00a0{\u0001\u0000\u0000"+
		"\u0000\u00a0~\u0001\u0000\u0000\u0000\u00a0\u0085\u0001\u0000\u0000\u0000"+
		"\u00a0\u008a\u0001\u0000\u0000\u0000\u00a0\u008d\u0001\u0000\u0000\u0000"+
		"\u00a0\u0092\u0001\u0000\u0000\u0000\u00a0\u0099\u0001\u0000\u0000\u0000"+
		"\u00a1\u000b\u0001\u0000\u0000\u0000\u00a2\u00a3\u0005:\u0000\u0000\u00a3"+
		"\u00a4\u0005;\u0000\u0000\u00a4\u00a5\u0005\b\u0000\u0000\u00a5\r\u0001"+
		"\u0000\u0000\u0000\u00a6\u00a8\u0005<\u0000\u0000\u00a7\u00a9\u0003D\""+
		"\u0000\u00a8\u00a7\u0001\u0000\u0000\u0000\u00a8\u00a9\u0001\u0000\u0000"+
		"\u0000\u00a9\u00aa\u0001\u0000\u0000\u0000\u00aa\u00ab\u0005=\u0000\u0000"+
		"\u00ab\u000f\u0001\u0000\u0000\u0000\u00ac\u00ad\u0005#\u0000\u0000\u00ad"+
		"\u00b6\u00058\u0000\u0000\u00ae\u00b3\u0003<\u001e\u0000\u00af\u00b0\u0005"+
		"A\u0000\u0000\u00b0\u00b2\u0003<\u001e\u0000\u00b1\u00af\u0001\u0000\u0000"+
		"\u0000\u00b2\u00b5\u0001\u0000\u0000\u0000\u00b3\u00b1\u0001\u0000\u0000"+
		"\u0000\u00b3\u00b4\u0001\u0000\u0000\u0000\u00b4\u00b7\u0001\u0000\u0000"+
		"\u0000\u00b5\u00b3\u0001\u0000\u0000\u0000\u00b6\u00ae\u0001\u0000\u0000"+
		"\u0000\u00b6\u00b7\u0001\u0000\u0000\u0000\u00b7\u00b8\u0001\u0000\u0000"+
		"\u0000\u00b8\u00bd\u00059\u0000\u0000\u00b9\u00bd\u0003<\u001e\u0000\u00ba"+
		"\u00bd\u0003\u0012\t\u0000\u00bb\u00bd\u0003\u0014\n\u0000\u00bc\u00ac"+
		"\u0001\u0000\u0000\u0000\u00bc\u00b9\u0001\u0000\u0000\u0000\u00bc\u00ba"+
		"\u0001\u0000\u0000\u0000\u00bc\u00bb\u0001\u0000\u0000\u0000\u00bd\u0011"+
		"\u0001\u0000\u0000\u0000\u00be\u00c3\u0003\u0016\u000b\u0000\u00bf\u00c3"+
		"\u0003\u001c\u000e\u0000\u00c0\u00c3\u0003 \u0010\u0000\u00c1\u00c3\u0003"+
		":\u001d\u0000\u00c2\u00be\u0001\u0000\u0000\u0000\u00c2\u00bf\u0001\u0000"+
		"\u0000\u0000\u00c2\u00c0\u0001\u0000\u0000\u0000\u00c2\u00c1\u0001\u0000"+
		"\u0000\u0000\u00c3\u0013\u0001\u0000\u0000\u0000\u00c4\u00cb\u0005\u0016"+
		"\u0000\u0000\u00c5\u00cb\u0005\u0017\u0000\u0000\u00c6\u00c8\u0005\u0018"+
		"\u0000\u0000\u00c7\u00c9\u0003<\u001e\u0000\u00c8\u00c7\u0001\u0000\u0000"+
		"\u0000\u00c8\u00c9\u0001\u0000\u0000\u0000\u00c9\u00cb\u0001\u0000\u0000"+
		"\u0000\u00ca\u00c4\u0001\u0000\u0000\u0000\u00ca\u00c5\u0001\u0000\u0000"+
		"\u0000\u00ca\u00c6\u0001\u0000\u0000\u0000\u00cb\u0015\u0001\u0000\u0000"+
		"\u0000\u00cc\u00cd\u0005\u0010\u0000\u0000\u00cd\u00ce\u0003<\u001e\u0000"+
		"\u00ce\u00d2\u0005<\u0000\u0000\u00cf\u00d1\u0003\b\u0004\u0000\u00d0"+
		"\u00cf\u0001\u0000\u0000\u0000\u00d1\u00d4\u0001\u0000\u0000\u0000\u00d2"+
		"\u00d0\u0001\u0000\u0000\u0000\u00d2\u00d3\u0001\u0000\u0000\u0000\u00d3"+
		"\u00d5\u0001\u0000\u0000\u0000\u00d4\u00d2\u0001\u0000\u0000\u0000\u00d5"+
		"\u00d9\u0005=\u0000\u0000\u00d6\u00d8\u0003\u0018\f\u0000\u00d7\u00d6"+
		"\u0001\u0000\u0000\u0000\u00d8\u00db\u0001\u0000\u0000\u0000\u00d9\u00d7"+
		"\u0001\u0000\u0000\u0000\u00d9\u00da\u0001\u0000\u0000\u0000\u00da\u00dd"+
		"\u0001\u0000\u0000\u0000\u00db\u00d9\u0001\u0000\u0000\u0000\u00dc\u00de"+
		"\u0003\u001a\r\u0000\u00dd\u00dc\u0001\u0000\u0000\u0000\u00dd\u00de\u0001"+
		"\u0000\u0000\u0000\u00de\u0017\u0001\u0000\u0000\u0000\u00df\u00e0\u0005"+
		"\u0011\u0000\u0000\u00e0\u00e1\u0005\u0010\u0000\u0000\u00e1\u00e2\u0003"+
		"<\u001e\u0000\u00e2\u00e6\u0005<\u0000\u0000\u00e3\u00e5\u0003\b\u0004"+
		"\u0000\u00e4\u00e3\u0001\u0000\u0000\u0000\u00e5\u00e8\u0001\u0000\u0000"+
		"\u0000\u00e6\u00e4\u0001\u0000\u0000\u0000\u00e6\u00e7\u0001\u0000\u0000"+
		"\u0000\u00e7\u00e9\u0001\u0000\u0000\u0000\u00e8\u00e6\u0001\u0000\u0000"+
		"\u0000\u00e9\u00ea\u0005=\u0000\u0000\u00ea\u0019\u0001\u0000\u0000\u0000"+
		"\u00eb\u00ec\u0005\u0011\u0000\u0000\u00ec\u00f0\u0005<\u0000\u0000\u00ed"+
		"\u00ef\u0003\b\u0004\u0000\u00ee\u00ed\u0001\u0000\u0000\u0000\u00ef\u00f2"+
		"\u0001\u0000\u0000\u0000\u00f0\u00ee\u0001\u0000\u0000\u0000\u00f0\u00f1"+
		"\u0001\u0000\u0000\u0000\u00f1\u00f3\u0001\u0000\u0000\u0000\u00f2\u00f0"+
		"\u0001\u0000\u0000\u0000\u00f3\u00f4\u0005=\u0000\u0000\u00f4\u001b\u0001"+
		"\u0000\u0000\u0000\u00f5\u00f6\u0005\u0012\u0000\u0000\u00f6\u00f7\u0003"+
		"\u001e\u000f\u0000\u00f7\u00f8\u0005>\u0000\u0000\u00f8\u00f9\u0003<\u001e"+
		"\u0000\u00f9\u00fb\u0005>\u0000\u0000\u00fa\u00fc\u0003\u0010\b\u0000"+
		"\u00fb\u00fa\u0001\u0000\u0000\u0000\u00fb\u00fc\u0001\u0000\u0000\u0000"+
		"\u00fc\u00fd\u0001\u0000\u0000\u0000\u00fd\u00fe\u0003\u0006\u0003\u0000"+
		"\u00fe\u0104\u0001\u0000\u0000\u0000\u00ff\u0100\u0005\u0012\u0000\u0000"+
		"\u0100\u0101\u0003<\u001e\u0000\u0101\u0102\u0003\u0006\u0003\u0000\u0102"+
		"\u0104\u0001\u0000\u0000\u0000\u0103\u00f5\u0001\u0000\u0000\u0000\u0103"+
		"\u00ff\u0001\u0000\u0000\u0000\u0104\u001d\u0001\u0000\u0000\u0000\u0105"+
		"\u0106\u0005$\u0000\u0000\u0106\u0107\u00057\u0000\u0000\u0107\u0108\u0003"+
		"<\u001e\u0000\u0108\u001f\u0001\u0000\u0000\u0000\u0109\u010a\u0005\u0013"+
		"\u0000\u0000\u010a\u010b\u0003<\u001e\u0000\u010b\u010f\u0005<\u0000\u0000"+
		"\u010c\u010e\u0003\"\u0011\u0000\u010d\u010c\u0001\u0000\u0000\u0000\u010e"+
		"\u0111\u0001\u0000\u0000\u0000\u010f\u010d\u0001\u0000\u0000\u0000\u010f"+
		"\u0110\u0001\u0000\u0000\u0000\u0110\u0113\u0001\u0000\u0000\u0000\u0111"+
		"\u010f\u0001\u0000\u0000\u0000\u0112\u0114\u0003$\u0012\u0000\u0113\u0112"+
		"\u0001\u0000\u0000\u0000\u0113\u0114\u0001\u0000\u0000\u0000\u0114\u0115"+
		"\u0001\u0000\u0000\u0000\u0115\u0116\u0005=\u0000\u0000\u0116!\u0001\u0000"+
		"\u0000\u0000\u0117\u0118\u0005\u0004\u0000\u0000\u0118\u0119\u0003<\u001e"+
		"\u0000\u0119\u011d\u0005?\u0000\u0000\u011a\u011c\u0003\b\u0004\u0000"+
		"\u011b\u011a\u0001\u0000\u0000\u0000\u011c\u011f\u0001\u0000\u0000\u0000"+
		"\u011d\u011b\u0001\u0000\u0000\u0000\u011d\u011e\u0001\u0000\u0000\u0000"+
		"\u011e#\u0001\u0000\u0000\u0000\u011f\u011d\u0001\u0000\u0000\u0000\u0120"+
		"\u0121\u0005\u0005\u0000\u0000\u0121\u0125\u0005?\u0000\u0000\u0122\u0124"+
		"\u0003\b\u0004\u0000\u0123\u0122\u0001\u0000\u0000\u0000\u0124\u0127\u0001"+
		"\u0000\u0000\u0000\u0125\u0123\u0001\u0000\u0000\u0000\u0125\u0126\u0001"+
		"\u0000\u0000\u0000\u0126%\u0001\u0000\u0000\u0000\u0127\u0125\u0001\u0000"+
		"\u0000\u0000\u0128\u0129\u0005\u0014\u0000\u0000\u0129\u0132\u00058\u0000"+
		"\u0000\u012a\u012f\u0003<\u001e\u0000\u012b\u012c\u0005A\u0000\u0000\u012c"+
		"\u012e\u0003<\u001e\u0000\u012d\u012b\u0001\u0000\u0000\u0000\u012e\u0131"+
		"\u0001\u0000\u0000\u0000\u012f\u012d\u0001\u0000\u0000\u0000\u012f\u0130"+
		"\u0001\u0000\u0000\u0000\u0130\u0133\u0001\u0000\u0000\u0000\u0131\u012f"+
		"\u0001\u0000\u0000\u0000\u0132\u012a\u0001\u0000\u0000\u0000\u0132\u0133"+
		"\u0001\u0000\u0000\u0000\u0133\u0134\u0001\u0000\u0000\u0000\u0134\u016a"+
		"\u00059\u0000\u0000\u0135\u0136\u0005\u0015\u0000\u0000\u0136\u013f\u0005"+
		"8\u0000\u0000\u0137\u013c\u0003<\u001e\u0000\u0138\u0139\u0005A\u0000"+
		"\u0000\u0139\u013b\u0003<\u001e\u0000\u013a\u0138\u0001\u0000\u0000\u0000"+
		"\u013b\u013e\u0001\u0000\u0000\u0000\u013c\u013a\u0001\u0000\u0000\u0000"+
		"\u013c\u013d\u0001\u0000\u0000\u0000\u013d\u0140\u0001\u0000\u0000\u0000"+
		"\u013e\u013c\u0001\u0000\u0000\u0000\u013f\u0137\u0001\u0000\u0000\u0000"+
		"\u013f\u0140\u0001\u0000\u0000\u0000\u0140\u0141\u0001\u0000\u0000\u0000"+
		"\u0141\u016a\u00059\u0000\u0000\u0142\u0143\u0005$\u0000\u0000\u0143\u014c"+
		"\u00058\u0000\u0000\u0144\u0149\u0003<\u001e\u0000\u0145\u0146\u0005A"+
		"\u0000\u0000\u0146\u0148\u0003<\u001e\u0000\u0147\u0145\u0001\u0000\u0000"+
		"\u0000\u0148\u014b\u0001\u0000\u0000\u0000\u0149\u0147\u0001\u0000\u0000"+
		"\u0000\u0149\u014a\u0001\u0000\u0000\u0000\u014a\u014d\u0001\u0000\u0000"+
		"\u0000\u014b\u0149\u0001\u0000\u0000\u0000\u014c\u0144\u0001\u0000\u0000"+
		"\u0000\u014c\u014d\u0001\u0000\u0000\u0000\u014d\u014e\u0001\u0000\u0000"+
		"\u0000\u014e\u016a\u00059\u0000\u0000\u014f\u0150\u0005\r\u0000\u0000"+
		"\u0150\u0159\u00058\u0000\u0000\u0151\u0156\u0003<\u001e\u0000\u0152\u0153"+
		"\u0005A\u0000\u0000\u0153\u0155\u0003<\u001e\u0000\u0154\u0152\u0001\u0000"+
		"\u0000\u0000\u0155\u0158\u0001\u0000\u0000\u0000\u0156\u0154\u0001\u0000"+
		"\u0000\u0000\u0156\u0157\u0001\u0000\u0000\u0000\u0157\u015a\u0001\u0000"+
		"\u0000\u0000\u0158\u0156\u0001\u0000\u0000\u0000\u0159\u0151\u0001\u0000"+
		"\u0000\u0000\u0159\u015a\u0001\u0000\u0000\u0000\u015a\u015b\u0001\u0000"+
		"\u0000\u0000\u015b\u016a\u00059\u0000\u0000\u015c\u015d\u0005\u000f\u0000"+
		"\u0000\u015d\u0166\u00058\u0000\u0000\u015e\u0163\u0003<\u001e\u0000\u015f"+
		"\u0160\u0005A\u0000\u0000\u0160\u0162\u0003<\u001e\u0000\u0161\u015f\u0001"+
		"\u0000\u0000\u0000\u0162\u0165\u0001\u0000\u0000\u0000\u0163\u0161\u0001"+
		"\u0000\u0000\u0000\u0163\u0164\u0001\u0000\u0000\u0000\u0164\u0167\u0001"+
		"\u0000\u0000\u0000\u0165\u0163\u0001\u0000\u0000\u0000\u0166\u015e\u0001"+
		"\u0000\u0000\u0000\u0166\u0167\u0001\u0000\u0000\u0000\u0167\u0168\u0001"+
		"\u0000\u0000\u0000\u0168\u016a\u00059\u0000\u0000\u0169\u0128\u0001\u0000"+
		"\u0000\u0000\u0169\u0135\u0001\u0000\u0000\u0000\u0169\u0142\u0001\u0000"+
		"\u0000\u0000\u0169\u014f\u0001\u0000\u0000\u0000\u0169\u015c\u0001\u0000"+
		"\u0000\u0000\u016a\'\u0001\u0000\u0000\u0000\u016b\u016c\u0005$\u0000"+
		"\u0000\u016c\u016e\u00058\u0000\u0000\u016d\u016f\u0003.\u0017\u0000\u016e"+
		"\u016d\u0001\u0000\u0000\u0000\u016e\u016f\u0001\u0000\u0000\u0000\u016f"+
		"\u0170\u0001\u0000\u0000\u0000\u0170\u0171\u00059\u0000\u0000\u0171)\u0001"+
		"\u0000\u0000\u0000\u0172\u0177\u0003,\u0016\u0000\u0173\u0174\u0005A\u0000"+
		"\u0000\u0174\u0176\u0003,\u0016\u0000\u0175\u0173\u0001\u0000\u0000\u0000"+
		"\u0176\u0179\u0001\u0000\u0000\u0000\u0177\u0175\u0001\u0000\u0000\u0000"+
		"\u0177\u0178\u0001\u0000\u0000\u0000\u0178+\u0001\u0000\u0000\u0000\u0179"+
		"\u0177\u0001\u0000\u0000\u0000\u017a\u017b\u0005$\u0000\u0000\u017b\u017c"+
		"\u0005\b\u0000\u0000\u017c-\u0001\u0000\u0000\u0000\u017d\u0182\u0003"+
		"<\u001e\u0000\u017e\u017f\u0005A\u0000\u0000\u017f\u0181\u0003<\u001e"+
		"\u0000\u0180\u017e\u0001\u0000\u0000\u0000\u0181\u0184\u0001\u0000\u0000"+
		"\u0000\u0182\u0180\u0001\u0000\u0000\u0000\u0182\u0183\u0001\u0000\u0000"+
		"\u0000\u0183/\u0001\u0000\u0000\u0000\u0184\u0182\u0001\u0000\u0000\u0000"+
		"\u0185\u0186\u0005\u0006\u0000\u0000\u0186\u0187\u0005$\u0000\u0000\u0187"+
		"\u0188\u0005<\u0000\u0000\u0188\u0189\u00032\u0019\u0000\u0189\u018a\u0005"+
		"=\u0000\u0000\u018a1\u0001\u0000\u0000\u0000\u018b\u018d\u00034\u001a"+
		"\u0000\u018c\u018b\u0001\u0000\u0000\u0000\u018d\u018e\u0001\u0000\u0000"+
		"\u0000\u018e\u018c\u0001\u0000\u0000\u0000\u018e\u018f\u0001\u0000\u0000"+
		"\u0000\u018f3\u0001\u0000\u0000\u0000\u0190\u0191\u0005$\u0000\u0000\u0191"+
		"\u0195\u0005\b\u0000\u0000\u0192\u0193\u0005$\u0000\u0000\u0193\u0195"+
		"\u0005$\u0000\u0000\u0194\u0190\u0001\u0000\u0000\u0000\u0194\u0192\u0001"+
		"\u0000\u0000\u0000\u01955\u0001\u0000\u0000\u0000\u0196\u019b\u00038\u001c"+
		"\u0000\u0197\u0198\u0005A\u0000\u0000\u0198\u019a\u00038\u001c\u0000\u0199"+
		"\u0197\u0001\u0000\u0000\u0000\u019a\u019d\u0001\u0000\u0000\u0000\u019b"+
		"\u0199\u0001\u0000\u0000\u0000\u019b\u019c\u0001\u0000\u0000\u0000\u019c"+
		"7\u0001\u0000\u0000\u0000\u019d\u019b\u0001\u0000\u0000\u0000\u019e\u019f"+
		"\u0005$\u0000\u0000\u019f\u01a0\u0005?\u0000\u0000\u01a0\u01a1\u0003<"+
		"\u001e\u0000\u01a19\u0001\u0000\u0000\u0000\u01a2\u01a3\u0005\u0007\u0000"+
		"\u0000\u01a3\u01a4\u00058\u0000\u0000\u01a4\u01a5\u0003<\u001e\u0000\u01a5"+
		"\u01a6\u00059\u0000\u0000\u01a6\u01aa\u0005:\u0000\u0000\u01a7\u01a9\u0003"+
		"\b\u0004\u0000\u01a8\u01a7\u0001\u0000\u0000\u0000\u01a9\u01ac\u0001\u0000"+
		"\u0000\u0000\u01aa\u01a8\u0001\u0000\u0000\u0000\u01aa\u01ab\u0001\u0000"+
		"\u0000\u0000\u01ab\u01ad\u0001\u0000\u0000\u0000\u01ac\u01aa\u0001\u0000"+
		"\u0000\u0000\u01ad\u01ae\u0005;\u0000\u0000\u01ae;\u0001\u0000\u0000\u0000"+
		"\u01af\u01b0\u0006\u001e\uffff\uffff\u0000\u01b0\u01b1\u0007\u0000\u0000"+
		"\u0000\u01b1\u01ee\u0003<\u001e\u0013\u01b2\u01ee\u0003B!\u0000\u01b3"+
		"\u01b4\u00058\u0000\u0000\u01b4\u01b5\u0003<\u001e\u0000\u01b5\u01b6\u0005"+
		"9\u0000\u0000\u01b6\u01ee\u0001\u0000\u0000\u0000\u01b7\u01b8\u0005:\u0000"+
		"\u0000\u01b8\u01b9\u0003<\u001e\u0000\u01b9\u01ba\u0005;\u0000\u0000\u01ba"+
		"\u01ee\u0001\u0000\u0000\u0000\u01bb\u01bc\u0005$\u0000\u0000\u01bc\u01bd"+
		"\u0005:\u0000\u0000\u01bd\u01be\u0003<\u001e\u0000\u01be\u01bf\u0005;"+
		"\u0000\u0000\u01bf\u01ee\u0001\u0000\u0000\u0000\u01c0\u01ee\u0003&\u0013"+
		"\u0000\u01c1\u01c2\u0005$\u0000\u0000\u01c2\u01c3\u0005@\u0000\u0000\u01c3"+
		"\u01c4\u0005$\u0000\u0000\u01c4\u01c5\u00057\u0000\u0000\u01c5\u01ee\u0003"+
		"<\u001e\n\u01c6\u01c7\u0005$\u0000\u0000\u01c7\u01c8\u0005<\u0000\u0000"+
		"\u01c8\u01c9\u00036\u001b\u0000\u01c9\u01ca\u0005=\u0000\u0000\u01ca\u01ee"+
		"\u0001\u0000\u0000\u0000\u01cb\u01ee\u0005$\u0000\u0000\u01cc\u01ee\u0003"+
		"F#\u0000\u01cd\u01ce\u0005$\u0000\u0000\u01ce\u01cf\u0005@\u0000\u0000"+
		"\u01cf\u01ee\u0005$\u0000\u0000\u01d0\u01d1\u0005$\u0000\u0000\u01d1\u01d2"+
		"\u0005@\u0000\u0000\u01d2\u01ee\u0003<\u001e\u0005\u01d3\u01d4\u0005$"+
		"\u0000\u0000\u01d4\u01d5\u0005\b\u0000\u0000\u01d5\u01d6\u00057\u0000"+
		"\u0000\u01d6\u01ee\u0003<\u001e\u0004\u01d7\u01d8\u0005$\u0000\u0000\u01d8"+
		"\u01d9\u0007\u0001\u0000\u0000\u01d9\u01ee\u0003<\u001e\u0003\u01da\u01db"+
		"\u0005\f\u0000\u0000\u01db\u01dc\u00058\u0000\u0000\u01dc\u01dd\u0005"+
		":\u0000\u0000\u01dd\u01e2\u0003<\u001e\u0000\u01de\u01df\u0005A\u0000"+
		"\u0000\u01df\u01e1\u0003<\u001e\u0000\u01e0\u01de\u0001\u0000\u0000\u0000"+
		"\u01e1\u01e4\u0001\u0000\u0000\u0000\u01e2\u01e0\u0001\u0000\u0000\u0000"+
		"\u01e2\u01e3\u0001\u0000\u0000\u0000\u01e3\u01e5\u0001\u0000\u0000\u0000"+
		"\u01e4\u01e2\u0001\u0000\u0000\u0000\u01e5\u01e6\u0005;\u0000\u0000\u01e6"+
		"\u01e7\u00059\u0000\u0000\u01e7\u01ee\u0001\u0000\u0000\u0000\u01e8\u01e9"+
		"\u0005\f\u0000\u0000\u01e9\u01ea\u00058\u0000\u0000\u01ea\u01eb\u0003"+
		"<\u001e\u0000\u01eb\u01ec\u00059\u0000\u0000\u01ec\u01ee\u0001\u0000\u0000"+
		"\u0000\u01ed\u01af\u0001\u0000\u0000\u0000\u01ed\u01b2\u0001\u0000\u0000"+
		"\u0000\u01ed\u01b3\u0001\u0000\u0000\u0000\u01ed\u01b7\u0001\u0000\u0000"+
		"\u0000\u01ed\u01bb\u0001\u0000\u0000\u0000\u01ed\u01c0\u0001\u0000\u0000"+
		"\u0000\u01ed\u01c1\u0001\u0000\u0000\u0000\u01ed\u01c6\u0001\u0000\u0000"+
		"\u0000\u01ed\u01cb\u0001\u0000\u0000\u0000\u01ed\u01cc\u0001\u0000\u0000"+
		"\u0000\u01ed\u01cd\u0001\u0000\u0000\u0000\u01ed\u01d0\u0001\u0000\u0000"+
		"\u0000\u01ed\u01d3\u0001\u0000\u0000\u0000\u01ed\u01d7\u0001\u0000\u0000"+
		"\u0000\u01ed\u01da\u0001\u0000\u0000\u0000\u01ed\u01e8\u0001\u0000\u0000"+
		"\u0000\u01ee\u0200\u0001\u0000\u0000\u0000\u01ef\u01f0\n\u0015\u0000\u0000"+
		"\u01f0\u01f1\u0007\u0002\u0000\u0000\u01f1\u01ff\u0003<\u001e\u0016\u01f2"+
		"\u01f3\n\u0014\u0000\u0000\u01f3\u01f4\u0007\u0003\u0000\u0000\u01f4\u01ff"+
		"\u0003<\u001e\u0015\u01f5\u01f6\n\u0012\u0000\u0000\u01f6\u01f7\u0007"+
		"\u0004\u0000\u0000\u01f7\u01ff\u0003<\u001e\u0013\u01f8\u01f9\n\u0011"+
		"\u0000\u0000\u01f9\u01fa\u0007\u0005\u0000\u0000\u01fa\u01ff\u0003<\u001e"+
		"\u0012\u01fb\u01fc\n\u0010\u0000\u0000\u01fc\u01fd\u0007\u0006\u0000\u0000"+
		"\u01fd\u01ff\u0003<\u001e\u0011\u01fe\u01ef\u0001\u0000\u0000\u0000\u01fe"+
		"\u01f2\u0001\u0000\u0000\u0000\u01fe\u01f5\u0001\u0000\u0000\u0000\u01fe"+
		"\u01f8\u0001\u0000\u0000\u0000\u01fe\u01fb\u0001\u0000\u0000\u0000\u01ff"+
		"\u0202\u0001\u0000\u0000\u0000\u0200\u01fe\u0001\u0000\u0000\u0000\u0200"+
		"\u0201\u0001\u0000\u0000\u0000\u0201=\u0001\u0000\u0000\u0000\u0202\u0200"+
		"\u0001\u0000\u0000\u0000\u0203\u0208\u0003<\u001e\u0000\u0204\u0205\u0005"+
		"A\u0000\u0000\u0205\u0207\u0003<\u001e\u0000\u0206\u0204\u0001\u0000\u0000"+
		"\u0000\u0207\u020a\u0001\u0000\u0000\u0000\u0208\u0206\u0001\u0000\u0000"+
		"\u0000\u0208\u0209\u0001\u0000\u0000\u0000\u0209?\u0001\u0000\u0000\u0000"+
		"\u020a\u0208\u0001\u0000\u0000\u0000\u020b\u020c\u0003B!\u0000\u020cA"+
		"\u0001\u0000\u0000\u0000\u020d\u0213\u0005\u001a\u0000\u0000\u020e\u0213"+
		"\u0005\u001b\u0000\u0000\u020f\u0213\u0005\u001c\u0000\u0000\u0210\u0213"+
		"\u0005\u0019\u0000\u0000\u0211\u0213\u0005\u001d\u0000\u0000\u0212\u020d"+
		"\u0001\u0000\u0000\u0000\u0212\u020e\u0001\u0000\u0000\u0000\u0212\u020f"+
		"\u0001\u0000\u0000\u0000\u0212\u0210\u0001\u0000\u0000\u0000\u0212\u0211"+
		"\u0001\u0000\u0000\u0000\u0213C\u0001\u0000\u0000\u0000\u0214\u0219\u0003"+
		"<\u001e\u0000\u0215\u0216\u0005A\u0000\u0000\u0216\u0218\u0003<\u001e"+
		"\u0000\u0217\u0215\u0001\u0000\u0000\u0000\u0218\u021b\u0001\u0000\u0000"+
		"\u0000\u0219\u0217\u0001\u0000\u0000\u0000\u0219\u021a\u0001\u0000\u0000"+
		"\u0000\u021aE\u0001\u0000\u0000\u0000\u021b\u0219\u0001\u0000\u0000\u0000"+
		"\u021c\u021d\u0005$\u0000\u0000\u021d\u0221\u0005%\u0000\u0000\u021e\u021f"+
		"\u0005$\u0000\u0000\u021f\u0221\u0005&\u0000\u0000\u0220\u021c\u0001\u0000"+
		"\u0000\u0000\u0220\u021e\u0001\u0000\u0000\u0000\u0221G\u0001\u0000\u0000"+
		"\u00005KZ^fpuy\u0090\u00a0\u00a8\u00b3\u00b6\u00bc\u00c2\u00c8\u00ca\u00d2"+
		"\u00d9\u00dd\u00e6\u00f0\u00fb\u0103\u010f\u0113\u011d\u0125\u012f\u0132"+
		"\u013c\u013f\u0149\u014c\u0156\u0159\u0163\u0166\u0169\u016e\u0177\u0182"+
		"\u018e\u0194\u019b\u01aa\u01e2\u01ed\u01fe\u0200\u0208\u0212\u0219\u0220";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}