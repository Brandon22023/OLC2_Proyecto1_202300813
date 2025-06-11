// Generated from /home/vboxuser/Documents/OLC2_Proyecto1_202300813/Backend/compiler/parser/Vlang.g4 by ANTLR 4.13.1
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
		ATOI=10, PARSEFLOAT=11, LEN=12, CAP=13, APPEND=14, IF=15, ELSE=16, FOR=17, 
		SWITCH=18, INDEXOF=19, JOIN=20, BREAK=21, CONTINUE=22, RETURN=23, BOOLEANO=24, 
		ENTERO=25, DECIMAL=26, CADENA=27, CARACTER=28, TIPO_ENTERO=29, TIPO_DECIMAL=30, 
		TIPO_CADENA=31, TIPO_BOOLEANO=32, TIPO_CHAR=33, PRINT=34, ID=35, INC=36, 
		DEC=37, SUMAIMPLICITA=38, RESTOIMPLICITO=39, PLUS=40, MINUS=41, MUL=42, 
		DIV=43, MOD=44, NOT=45, OR=46, AND=47, EQ=48, NEQ=49, LE=50, GE=51, LT=52, 
		GT=53, ASSIGN=54, LPAREN=55, RPAREN=56, LBRACK=57, RBRACK=58, LBRACE=59, 
		RBRACE=60, SEMICOLON=61, COLON=62, DOT=63, COMMA=64, WS=65, LINE_COMMENT=66, 
		BLOCK_COMMENT=67;
	public static final int
		RULE_programa = 0, RULE_funcMain = 1, RULE_block = 2, RULE_declaraciones = 3, 
		RULE_varDcl = 4, RULE_sliceTipo = 5, RULE_sliceInit = 6, RULE_stmt = 7, 
		RULE_sentencias_control = 8, RULE_sentencias_transferencia = 9, RULE_ifDcl = 10, 
		RULE_elseIfDcl = 11, RULE_elseCondicional = 12, RULE_forDcl = 13, RULE_asignacion = 14, 
		RULE_switchDcl = 15, RULE_caseBlock = 16, RULE_defaultBlock = 17, RULE_llamadaFuncion = 18, 
		RULE_funcDcl = 19, RULE_funcCall = 20, RULE_parametrosFormales = 21, RULE_parametro = 22, 
		RULE_parametrosReales = 23, RULE_whileDcl = 24, RULE_expresion = 25, RULE_parametros = 26, 
		RULE_valores = 27, RULE_valor = 28, RULE_listaExpresiones = 29, RULE_incredecre = 30;
	private static String[] makeRuleNames() {
		return new String[] {
			"programa", "funcMain", "block", "declaraciones", "varDcl", "sliceTipo", 
			"sliceInit", "stmt", "sentencias_control", "sentencias_transferencia", 
			"ifDcl", "elseIfDcl", "elseCondicional", "forDcl", "asignacion", "switchDcl", 
			"caseBlock", "defaultBlock", "llamadaFuncion", "funcDcl", "funcCall", 
			"parametrosFormales", "parametro", "parametrosReales", "whileDcl", "expresion", 
			"parametros", "valores", "valor", "listaExpresiones", "incredecre"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'fn'", "'main'", "'mut'", "'case'", "'default'", "'func'", "'while'", 
			null, null, "'Atoi'", "'parseFloat'", "'len'", "'cap'", "'append'", "'if'", 
			"'else'", "'for'", "'switch'", "'indexOf'", "'join'", "'break'", "'continue'", 
			"'return'", null, null, null, null, null, "'int'", "'float64'", "'string'", 
			"'bool'", "'rune'", "'print'", null, "'++'", "'--'", "'+='", "'-='", 
			"'+'", "'-'", "'*'", "'/'", "'%'", "'!'", "'||'", "'&&'", "'=='", "'!='", 
			"'<='", "'>='", "'<'", "'>'", "'='", "'('", "')'", "'['", "']'", "'{'", 
			"'}'", "';'", "':'", "'.'", "','"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, "TIPO", "CASTEOS", "ATOI", 
			"PARSEFLOAT", "LEN", "CAP", "APPEND", "IF", "ELSE", "FOR", "SWITCH", 
			"INDEXOF", "JOIN", "BREAK", "CONTINUE", "RETURN", "BOOLEANO", "ENTERO", 
			"DECIMAL", "CADENA", "CARACTER", "TIPO_ENTERO", "TIPO_DECIMAL", "TIPO_CADENA", 
			"TIPO_BOOLEANO", "TIPO_CHAR", "PRINT", "ID", "INC", "DEC", "SUMAIMPLICITA", 
			"RESTOIMPLICITO", "PLUS", "MINUS", "MUL", "DIV", "MOD", "NOT", "OR", 
			"AND", "EQ", "NEQ", "LE", "GE", "LT", "GT", "ASSIGN", "LPAREN", "RPAREN", 
			"LBRACK", "RBRACK", "LBRACE", "RBRACE", "SEMICOLON", "COLON", "DOT", 
			"COMMA", "WS", "LINE_COMMENT", "BLOCK_COMMENT"
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
		public List<FuncDclContext> funcDcl() {
			return getRuleContexts(FuncDclContext.class);
		}
		public FuncDclContext funcDcl(int i) {
			return getRuleContext(FuncDclContext.class,i);
		}
		public List<FuncMainContext> funcMain() {
			return getRuleContexts(FuncMainContext.class);
		}
		public FuncMainContext funcMain(int i) {
			return getRuleContext(FuncMainContext.class,i);
		}
		public ProgramaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_programa; }
	}

	public final ProgramaContext programa() throws RecognitionException {
		ProgramaContext _localctx = new ProgramaContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_programa);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(66);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__0 || _la==T__5) {
				{
				setState(64);
				_errHandler.sync(this);
				switch (_input.LA(1)) {
				case T__5:
					{
					setState(62);
					funcDcl();
					}
					break;
				case T__0:
					{
					setState(63);
					funcMain();
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				}
				setState(68);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(69);
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
	}

	public final FuncMainContext funcMain() throws RecognitionException {
		FuncMainContext _localctx = new FuncMainContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_funcMain);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(71);
			match(T__0);
			setState(72);
			match(T__1);
			setState(73);
			match(LPAREN);
			setState(74);
			match(RPAREN);
			setState(75);
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
	}

	public final BlockContext block() throws RecognitionException {
		BlockContext _localctx = new BlockContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_block);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(77);
			match(LBRACE);
			setState(81);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 180181420566565064L) != 0)) {
				{
				{
				setState(78);
				declaraciones();
				}
				}
				setState(83);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(84);
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
		public DeclaracionesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaraciones; }
	}

	public final DeclaracionesContext declaraciones() throws RecognitionException {
		DeclaracionesContext _localctx = new DeclaracionesContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_declaraciones);
		try {
			setState(89);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(86);
				varDcl();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(87);
				stmt();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(88);
				funcDcl();
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
	public static class SliceEmptyDeclarationContext extends VarDclContext {
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public SliceTipoContext sliceTipo() {
			return getRuleContext(SliceTipoContext.class,0);
		}
		public SliceEmptyDeclarationContext(VarDclContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SliceAssignmentContext extends VarDclContext {
		public List<TerminalNode> ID() { return getTokens(VlangParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(VlangParser.ID, i);
		}
		public TerminalNode ASSIGN() { return getToken(VlangParser.ASSIGN, 0); }
		public SliceAssignmentContext(VarDclContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class VariableDeclarationImmutableContext extends VarDclContext {
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public TerminalNode ASSIGN() { return getToken(VlangParser.ASSIGN, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public VariableDeclarationImmutableContext(VarDclContext ctx) { copyFrom(ctx); }
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
	}

	public final VarDclContext varDcl() throws RecognitionException {
		VarDclContext _localctx = new VarDclContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_varDcl);
		int _la;
		try {
			setState(130);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
			case 1:
				_localctx = new VariableDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(91);
				match(T__2);
				setState(92);
				match(ID);
				setState(94);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==TIPO) {
					{
					setState(93);
					match(TIPO);
					}
				}

				setState(98);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==ASSIGN) {
					{
					setState(96);
					match(ASSIGN);
					setState(97);
					expresion(0);
					}
				}

				}
				break;
			case 2:
				_localctx = new SliceEmptyDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(100);
				match(T__2);
				setState(101);
				match(ID);
				setState(102);
				sliceTipo();
				}
				break;
			case 3:
				_localctx = new SliceInitDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(103);
				match(ID);
				setState(104);
				match(ASSIGN);
				setState(105);
				sliceTipo();
				setState(106);
				sliceInit();
				}
				break;
			case 4:
				_localctx = new SliceAssignmentContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(108);
				match(ID);
				setState(109);
				match(ASSIGN);
				setState(110);
				match(ID);
				}
				break;
			case 5:
				_localctx = new VariableDeclarationImmutableContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(111);
				match(ID);
				setState(114);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==ASSIGN) {
					{
					setState(112);
					match(ASSIGN);
					setState(113);
					expresion(0);
					}
				}

				}
				break;
			case 6:
				_localctx = new VariableCastDeclarationContext(_localctx);
				enterOuterAlt(_localctx, 6);
				{
				setState(116);
				match(ID);
				setState(117);
				match(ASSIGN);
				setState(118);
				match(CASTEOS);
				setState(119);
				match(LPAREN);
				setState(120);
				expresion(0);
				setState(121);
				match(RPAREN);
				}
				break;
			case 7:
				_localctx = new SliceAssignmentIndexContext(_localctx);
				enterOuterAlt(_localctx, 7);
				{
				setState(123);
				match(ID);
				setState(124);
				match(LBRACK);
				setState(125);
				expresion(0);
				setState(126);
				match(RBRACK);
				setState(127);
				match(ASSIGN);
				setState(128);
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
	}

	public final SliceTipoContext sliceTipo() throws RecognitionException {
		SliceTipoContext _localctx = new SliceTipoContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_sliceTipo);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(132);
			match(LBRACK);
			setState(133);
			match(RBRACK);
			setState(134);
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
	}

	public final SliceInitContext sliceInit() throws RecognitionException {
		SliceInitContext _localctx = new SliceInitContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_sliceInit);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(136);
			match(LBRACE);
			setState(138);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 180181403371589632L) != 0)) {
				{
				setState(137);
				listaExpresiones();
				}
			}

			setState(140);
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
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ControlStatementContext extends StmtContext {
		public Sentencias_controlContext sentencias_control() {
			return getRuleContext(Sentencias_controlContext.class,0);
		}
		public ControlStatementContext(StmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExpresionStatementContext extends StmtContext {
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public ExpresionStatementContext(StmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class TransfersentenceContext extends StmtContext {
		public Sentencias_transferenciaContext sentencias_transferencia() {
			return getRuleContext(Sentencias_transferenciaContext.class,0);
		}
		public TransfersentenceContext(StmtContext ctx) { copyFrom(ctx); }
	}

	public final StmtContext stmt() throws RecognitionException {
		StmtContext _localctx = new StmtContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_stmt);
		int _la;
		try {
			setState(158);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case PRINT:
				_localctx = new PrintStatementContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(142);
				match(PRINT);
				setState(143);
				match(LPAREN);
				setState(152);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 180181403371589632L) != 0)) {
					{
					setState(144);
					expresion(0);
					setState(149);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==COMMA) {
						{
						{
						setState(145);
						match(COMMA);
						setState(146);
						expresion(0);
						}
						}
						setState(151);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					}
				}

				setState(154);
				match(RPAREN);
				}
				break;
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
				setState(155);
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
				setState(156);
				sentencias_control();
				}
				break;
			case BREAK:
			case CONTINUE:
			case RETURN:
				_localctx = new TransfersentenceContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(157);
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
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Switch_contextContext extends Sentencias_controlContext {
		public SwitchDclContext switchDcl() {
			return getRuleContext(SwitchDclContext.class,0);
		}
		public Switch_contextContext(Sentencias_controlContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class If_contextContext extends Sentencias_controlContext {
		public IfDclContext ifDcl() {
			return getRuleContext(IfDclContext.class,0);
		}
		public If_contextContext(Sentencias_controlContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class For_contextContext extends Sentencias_controlContext {
		public ForDclContext forDcl() {
			return getRuleContext(ForDclContext.class,0);
		}
		public For_contextContext(Sentencias_controlContext ctx) { copyFrom(ctx); }
	}

	public final Sentencias_controlContext sentencias_control() throws RecognitionException {
		Sentencias_controlContext _localctx = new Sentencias_controlContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_sentencias_control);
		try {
			setState(164);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case IF:
				_localctx = new If_contextContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(160);
				ifDcl();
				}
				break;
			case FOR:
				_localctx = new For_contextContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(161);
				forDcl();
				}
				break;
			case SWITCH:
				_localctx = new Switch_contextContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(162);
				switchDcl();
				}
				break;
			case T__6:
				_localctx = new While_contextContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(163);
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
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ContinueStatementContext extends Sentencias_transferenciaContext {
		public TerminalNode CONTINUE() { return getToken(VlangParser.CONTINUE, 0); }
		public ContinueStatementContext(Sentencias_transferenciaContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ReturnStatementContext extends Sentencias_transferenciaContext {
		public TerminalNode RETURN() { return getToken(VlangParser.RETURN, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public ReturnStatementContext(Sentencias_transferenciaContext ctx) { copyFrom(ctx); }
	}

	public final Sentencias_transferenciaContext sentencias_transferencia() throws RecognitionException {
		Sentencias_transferenciaContext _localctx = new Sentencias_transferenciaContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_sentencias_transferencia);
		try {
			setState(172);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case BREAK:
				_localctx = new BreakStatementContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(166);
				match(BREAK);
				}
				break;
			case CONTINUE:
				_localctx = new ContinueStatementContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(167);
				match(CONTINUE);
				}
				break;
			case RETURN:
				_localctx = new ReturnStatementContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(168);
				match(RETURN);
				setState(170);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
				case 1:
					{
					setState(169);
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
	}

	public final IfDclContext ifDcl() throws RecognitionException {
		IfDclContext _localctx = new IfDclContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_ifDcl);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(174);
			match(IF);
			setState(175);
			expresion(0);
			setState(176);
			match(LBRACE);
			setState(180);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 180181420566565064L) != 0)) {
				{
				{
				setState(177);
				declaraciones();
				}
				}
				setState(182);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(183);
			match(RBRACE);
			setState(187);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,16,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(184);
					elseIfDcl();
					}
					} 
				}
				setState(189);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,16,_ctx);
			}
			setState(191);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ELSE) {
				{
				setState(190);
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
	}

	public final ElseIfDclContext elseIfDcl() throws RecognitionException {
		ElseIfDclContext _localctx = new ElseIfDclContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_elseIfDcl);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(193);
			match(ELSE);
			setState(194);
			match(IF);
			setState(195);
			expresion(0);
			setState(196);
			match(LBRACE);
			setState(200);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 180181420566565064L) != 0)) {
				{
				{
				setState(197);
				declaraciones();
				}
				}
				setState(202);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(203);
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
	}

	public final ElseCondicionalContext elseCondicional() throws RecognitionException {
		ElseCondicionalContext _localctx = new ElseCondicionalContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_elseCondicional);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(205);
			match(ELSE);
			setState(206);
			match(LBRACE);
			setState(210);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 180181420566565064L) != 0)) {
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
	}

	public final ForDclContext forDcl() throws RecognitionException {
		ForDclContext _localctx = new ForDclContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_forDcl);
		int _la;
		try {
			setState(229);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,21,_ctx) ) {
			case 1:
				_localctx = new ForClasicoContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(215);
				match(FOR);
				setState(216);
				asignacion();
				setState(217);
				match(SEMICOLON);
				setState(218);
				expresion(0);
				setState(219);
				match(SEMICOLON);
				setState(221);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 180181420566564992L) != 0)) {
					{
					setState(220);
					stmt();
					}
				}

				setState(223);
				block();
				}
				break;
			case 2:
				_localctx = new ForCondicionUnicaContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(225);
				match(FOR);
				setState(226);
				expresion(0);
				setState(227);
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
	}

	public final AsignacionContext asignacion() throws RecognitionException {
		AsignacionContext _localctx = new AsignacionContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_asignacion);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(231);
			match(ID);
			setState(232);
			match(ASSIGN);
			setState(233);
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
	}

	public final SwitchDclContext switchDcl() throws RecognitionException {
		SwitchDclContext _localctx = new SwitchDclContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_switchDcl);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(235);
			match(SWITCH);
			setState(236);
			expresion(0);
			setState(237);
			match(LBRACE);
			setState(241);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__3) {
				{
				{
				setState(238);
				caseBlock();
				}
				}
				setState(243);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(245);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__4) {
				{
				setState(244);
				defaultBlock();
				}
			}

			setState(247);
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
	}

	public final CaseBlockContext caseBlock() throws RecognitionException {
		CaseBlockContext _localctx = new CaseBlockContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_caseBlock);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(249);
			match(T__3);
			setState(250);
			expresion(0);
			setState(251);
			match(COLON);
			setState(255);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 180181420566565064L) != 0)) {
				{
				{
				setState(252);
				declaraciones();
				}
				}
				setState(257);
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
	}

	public final DefaultBlockContext defaultBlock() throws RecognitionException {
		DefaultBlockContext _localctx = new DefaultBlockContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_defaultBlock);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(258);
			match(T__4);
			setState(259);
			match(COLON);
			setState(263);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 180181420566565064L) != 0)) {
				{
				{
				setState(260);
				declaraciones();
				}
				}
				setState(265);
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
	}

	public final LlamadaFuncionContext llamadaFuncion() throws RecognitionException {
		LlamadaFuncionContext _localctx = new LlamadaFuncionContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_llamadaFuncion);
		int _la;
		try {
			setState(331);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INDEXOF:
				enterOuterAlt(_localctx, 1);
				{
				setState(266);
				match(INDEXOF);
				setState(267);
				match(LPAREN);
				setState(276);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 180181403371589632L) != 0)) {
					{
					setState(268);
					expresion(0);
					setState(273);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==COMMA) {
						{
						{
						setState(269);
						match(COMMA);
						setState(270);
						expresion(0);
						}
						}
						setState(275);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					}
				}

				setState(278);
				match(RPAREN);
				}
				break;
			case JOIN:
				enterOuterAlt(_localctx, 2);
				{
				setState(279);
				match(JOIN);
				setState(280);
				match(LPAREN);
				setState(289);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 180181403371589632L) != 0)) {
					{
					setState(281);
					expresion(0);
					setState(286);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==COMMA) {
						{
						{
						setState(282);
						match(COMMA);
						setState(283);
						expresion(0);
						}
						}
						setState(288);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					}
				}

				setState(291);
				match(RPAREN);
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 3);
				{
				setState(292);
				match(ID);
				setState(293);
				match(LPAREN);
				setState(302);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 180181403371589632L) != 0)) {
					{
					setState(294);
					expresion(0);
					setState(299);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==COMMA) {
						{
						{
						setState(295);
						match(COMMA);
						setState(296);
						expresion(0);
						}
						}
						setState(301);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					}
				}

				setState(304);
				match(RPAREN);
				}
				break;
			case LEN:
				enterOuterAlt(_localctx, 4);
				{
				setState(305);
				match(LEN);
				setState(306);
				match(LPAREN);
				setState(315);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 180181403371589632L) != 0)) {
					{
					setState(307);
					expresion(0);
					setState(312);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==COMMA) {
						{
						{
						setState(308);
						match(COMMA);
						setState(309);
						expresion(0);
						}
						}
						setState(314);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					}
				}

				setState(317);
				match(RPAREN);
				}
				break;
			case APPEND:
				enterOuterAlt(_localctx, 5);
				{
				setState(318);
				match(APPEND);
				setState(319);
				match(LPAREN);
				setState(328);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 180181403371589632L) != 0)) {
					{
					setState(320);
					expresion(0);
					setState(325);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==COMMA) {
						{
						{
						setState(321);
						match(COMMA);
						setState(322);
						expresion(0);
						}
						}
						setState(327);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					}
				}

				setState(330);
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
	}

	public final FuncDclContext funcDcl() throws RecognitionException {
		FuncDclContext _localctx = new FuncDclContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_funcDcl);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(333);
			match(T__5);
			setState(334);
			match(ID);
			setState(335);
			match(LPAREN);
			setState(337);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ID) {
				{
				setState(336);
				parametrosFormales();
				}
			}

			setState(339);
			match(RPAREN);
			setState(341);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==TIPO) {
				{
				setState(340);
				match(TIPO);
				}
			}

			setState(343);
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
	}

	public final FuncCallContext funcCall() throws RecognitionException {
		FuncCallContext _localctx = new FuncCallContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_funcCall);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(345);
			match(ID);
			setState(346);
			match(LPAREN);
			setState(348);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 180181403371589632L) != 0)) {
				{
				setState(347);
				parametrosReales();
				}
			}

			setState(350);
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
	}

	public final ParametrosFormalesContext parametrosFormales() throws RecognitionException {
		ParametrosFormalesContext _localctx = new ParametrosFormalesContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_parametrosFormales);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(352);
			parametro();
			setState(357);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(353);
				match(COMMA);
				setState(354);
				parametro();
				}
				}
				setState(359);
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
	}

	public final ParametroContext parametro() throws RecognitionException {
		ParametroContext _localctx = new ParametroContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_parametro);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(360);
			match(ID);
			setState(361);
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
	}

	public final ParametrosRealesContext parametrosReales() throws RecognitionException {
		ParametrosRealesContext _localctx = new ParametrosRealesContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_parametrosReales);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(363);
			expresion(0);
			setState(368);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(364);
				match(COMMA);
				setState(365);
				expresion(0);
				}
				}
				setState(370);
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
	}

	public final WhileDclContext whileDcl() throws RecognitionException {
		WhileDclContext _localctx = new WhileDclContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_whileDcl);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(371);
			match(T__6);
			setState(372);
			match(LPAREN);
			setState(373);
			expresion(0);
			setState(374);
			match(RPAREN);
			setState(375);
			match(LBRACK);
			setState(379);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 180181420566565064L) != 0)) {
				{
				{
				setState(376);
				declaraciones();
				}
				}
				setState(381);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(382);
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
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IncredecrContext extends ExpresionContext {
		public IncredecreContext incredecre() {
			return getRuleContext(IncredecreContext.class,0);
		}
		public IncredecrContext(ExpresionContext ctx) { copyFrom(ctx); }
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
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ValorexprContext extends ExpresionContext {
		public ValorContext valor() {
			return getRuleContext(ValorContext.class,0);
		}
		public ValorexprContext(ExpresionContext ctx) { copyFrom(ctx); }
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
	}
	@SuppressWarnings("CheckReturnValue")
	public static class LlamadaFuncionExprContext extends ExpresionContext {
		public LlamadaFuncionContext llamadaFuncion() {
			return getRuleContext(LlamadaFuncionContext.class,0);
		}
		public LlamadaFuncionExprContext(ExpresionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExpdotexpContext extends ExpresionContext {
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public TerminalNode DOT() { return getToken(VlangParser.DOT, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public ExpdotexpContext(ExpresionContext ctx) { copyFrom(ctx); }
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
	}
	@SuppressWarnings("CheckReturnValue")
	public static class CorchetesexpreContext extends ExpresionContext {
		public TerminalNode LBRACK() { return getToken(VlangParser.LBRACK, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode RBRACK() { return getToken(VlangParser.RBRACK, 0); }
		public CorchetesexpreContext(ExpresionContext ctx) { copyFrom(ctx); }
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
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ParentesisexpreContext extends ExpresionContext {
		public TerminalNode LPAREN() { return getToken(VlangParser.LPAREN, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(VlangParser.RPAREN, 0); }
		public ParentesisexpreContext(ExpresionContext ctx) { copyFrom(ctx); }
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
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IdContext extends ExpresionContext {
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public IdContext(ExpresionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class Expdotexp1Context extends ExpresionContext {
		public List<TerminalNode> ID() { return getTokens(VlangParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(VlangParser.ID, i);
		}
		public TerminalNode DOT() { return getToken(VlangParser.DOT, 0); }
		public Expdotexp1Context(ExpresionContext ctx) { copyFrom(ctx); }
	}

	public final ExpresionContext expresion() throws RecognitionException {
		return expresion(0);
	}

	private ExpresionContext expresion(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExpresionContext _localctx = new ExpresionContext(_ctx, _parentState);
		ExpresionContext _prevctx = _localctx;
		int _startState = 50;
		enterRecursionRule(_localctx, 50, RULE_expresion, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(417);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,43,_ctx) ) {
			case 1:
				{
				_localctx = new UnarioContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(385);
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
				setState(386);
				expresion(12);
				}
				break;
			case 2:
				{
				_localctx = new ValorexprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(387);
				valor();
				}
				break;
			case 3:
				{
				_localctx = new ParentesisexpreContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(388);
				match(LPAREN);
				setState(389);
				expresion(0);
				setState(390);
				match(RPAREN);
				}
				break;
			case 4:
				{
				_localctx = new CorchetesexpreContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(392);
				match(LBRACK);
				setState(393);
				expresion(0);
				setState(394);
				match(RBRACK);
				}
				break;
			case 5:
				{
				_localctx = new PARAPRINTSLICEContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(396);
				match(ID);
				setState(397);
				match(LBRACK);
				setState(398);
				expresion(0);
				setState(399);
				match(RBRACK);
				}
				break;
			case 6:
				{
				_localctx = new LlamadaFuncionExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(401);
				llamadaFuncion();
				}
				break;
			case 7:
				{
				_localctx = new IdContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(402);
				match(ID);
				}
				break;
			case 8:
				{
				_localctx = new IncredecrContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(403);
				incredecre();
				}
				break;
			case 9:
				{
				_localctx = new Expdotexp1Context(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(404);
				match(ID);
				setState(405);
				match(DOT);
				setState(406);
				match(ID);
				}
				break;
			case 10:
				{
				_localctx = new ExpdotexpContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(407);
				match(ID);
				setState(408);
				match(DOT);
				setState(409);
				expresion(3);
				}
				break;
			case 11:
				{
				_localctx = new AsignacionLUEGOContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(410);
				match(ID);
				setState(411);
				match(TIPO);
				setState(412);
				match(ASSIGN);
				setState(413);
				expresion(2);
				}
				break;
			case 12:
				{
				_localctx = new IMCPLICITContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(414);
				match(ID);
				setState(415);
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
				setState(416);
				expresion(1);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(436);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,45,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(434);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,44,_ctx) ) {
					case 1:
						{
						_localctx = new RelacionalesContext(new ExpresionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(419);
						if (!(precpred(_ctx, 17))) throw new FailedPredicateException(this, "precpred(_ctx, 17)");
						setState(420);
						((RelacionalesContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 16888498602639360L) != 0)) ) {
							((RelacionalesContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(421);
						expresion(18);
						}
						break;
					case 2:
						{
						_localctx = new IgualdadContext(new ExpresionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(422);
						if (!(precpred(_ctx, 16))) throw new FailedPredicateException(this, "precpred(_ctx, 16)");
						setState(423);
						((IgualdadContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==EQ || _la==NEQ) ) {
							((IgualdadContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(424);
						expresion(17);
						}
						break;
					case 3:
						{
						_localctx = new OPERADORESLOGICOSContext(new ExpresionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(425);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(426);
						((OPERADORESLOGICOSContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==OR || _la==AND) ) {
							((OPERADORESLOGICOSContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(427);
						expresion(16);
						}
						break;
					case 4:
						{
						_localctx = new MultdivmodContext(new ExpresionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(428);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(429);
						((MultdivmodContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 30786325577728L) != 0)) ) {
							((MultdivmodContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(430);
						expresion(15);
						}
						break;
					case 5:
						{
						_localctx = new SumresContext(new ExpresionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(431);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(432);
						((SumresContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==PLUS || _la==MINUS) ) {
							((SumresContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(433);
						expresion(14);
						}
						break;
					}
					} 
				}
				setState(438);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,45,_ctx);
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
	}

	public final ParametrosContext parametros() throws RecognitionException {
		ParametrosContext _localctx = new ParametrosContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_parametros);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(439);
			expresion(0);
			setState(444);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(440);
				match(COMMA);
				setState(441);
				expresion(0);
				}
				}
				setState(446);
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
	}

	public final ValoresContext valores() throws RecognitionException {
		ValoresContext _localctx = new ValoresContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_valores);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(447);
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
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ValorEnteroContext extends ValorContext {
		public TerminalNode ENTERO() { return getToken(VlangParser.ENTERO, 0); }
		public ValorEnteroContext(ValorContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ValorBooleanoContext extends ValorContext {
		public TerminalNode BOOLEANO() { return getToken(VlangParser.BOOLEANO, 0); }
		public ValorBooleanoContext(ValorContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ValorCaracterContext extends ValorContext {
		public TerminalNode CARACTER() { return getToken(VlangParser.CARACTER, 0); }
		public ValorCaracterContext(ValorContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ValorCadenaContext extends ValorContext {
		public TerminalNode CADENA() { return getToken(VlangParser.CADENA, 0); }
		public ValorCadenaContext(ValorContext ctx) { copyFrom(ctx); }
	}

	public final ValorContext valor() throws RecognitionException {
		ValorContext _localctx = new ValorContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_valor);
		try {
			setState(454);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case ENTERO:
				_localctx = new ValorEnteroContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(449);
				match(ENTERO);
				}
				break;
			case DECIMAL:
				_localctx = new ValorDecimalContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(450);
				match(DECIMAL);
				}
				break;
			case CADENA:
				_localctx = new ValorCadenaContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(451);
				match(CADENA);
				}
				break;
			case BOOLEANO:
				_localctx = new ValorBooleanoContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(452);
				match(BOOLEANO);
				}
				break;
			case CARACTER:
				_localctx = new ValorCaracterContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(453);
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
	}

	public final ListaExpresionesContext listaExpresiones() throws RecognitionException {
		ListaExpresionesContext _localctx = new ListaExpresionesContext(_ctx, getState());
		enterRule(_localctx, 58, RULE_listaExpresiones);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(456);
			expresion(0);
			setState(461);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(457);
				match(COMMA);
				setState(458);
				expresion(0);
				}
				}
				setState(463);
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
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DecrementoContext extends IncredecreContext {
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public TerminalNode DEC() { return getToken(VlangParser.DEC, 0); }
		public DecrementoContext(IncredecreContext ctx) { copyFrom(ctx); }
	}

	public final IncredecreContext incredecre() throws RecognitionException {
		IncredecreContext _localctx = new IncredecreContext(_ctx, getState());
		enterRule(_localctx, 60, RULE_incredecre);
		try {
			setState(468);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,49,_ctx) ) {
			case 1:
				_localctx = new IncrementoContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(464);
				match(ID);
				setState(465);
				match(INC);
				}
				break;
			case 2:
				_localctx = new DecrementoContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(466);
				match(ID);
				setState(467);
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
		case 25:
			return expresion_sempred((ExpresionContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expresion_sempred(ExpresionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 17);
		case 1:
			return precpred(_ctx, 16);
		case 2:
			return precpred(_ctx, 15);
		case 3:
			return precpred(_ctx, 14);
		case 4:
			return precpred(_ctx, 13);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001C\u01d7\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002\u0012\u0007\u0012"+
		"\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0002\u0015\u0007\u0015"+
		"\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017\u0002\u0018\u0007\u0018"+
		"\u0002\u0019\u0007\u0019\u0002\u001a\u0007\u001a\u0002\u001b\u0007\u001b"+
		"\u0002\u001c\u0007\u001c\u0002\u001d\u0007\u001d\u0002\u001e\u0007\u001e"+
		"\u0001\u0000\u0001\u0000\u0005\u0000A\b\u0000\n\u0000\f\u0000D\t\u0000"+
		"\u0001\u0000\u0001\u0000\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0002\u0001\u0002\u0005\u0002P\b\u0002"+
		"\n\u0002\f\u0002S\t\u0002\u0001\u0002\u0001\u0002\u0001\u0003\u0001\u0003"+
		"\u0001\u0003\u0003\u0003Z\b\u0003\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0003\u0004_\b\u0004\u0001\u0004\u0001\u0004\u0003\u0004c\b\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0003\u0004s\b\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0003"+
		"\u0004\u0083\b\u0004\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0006\u0001\u0006\u0003\u0006\u008b\b\u0006\u0001\u0006\u0001\u0006\u0001"+
		"\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0005\u0007\u0094"+
		"\b\u0007\n\u0007\f\u0007\u0097\t\u0007\u0003\u0007\u0099\b\u0007\u0001"+
		"\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0003\u0007\u009f\b\u0007\u0001"+
		"\b\u0001\b\u0001\b\u0001\b\u0003\b\u00a5\b\b\u0001\t\u0001\t\u0001\t\u0001"+
		"\t\u0003\t\u00ab\b\t\u0003\t\u00ad\b\t\u0001\n\u0001\n\u0001\n\u0001\n"+
		"\u0005\n\u00b3\b\n\n\n\f\n\u00b6\t\n\u0001\n\u0001\n\u0005\n\u00ba\b\n"+
		"\n\n\f\n\u00bd\t\n\u0001\n\u0003\n\u00c0\b\n\u0001\u000b\u0001\u000b\u0001"+
		"\u000b\u0001\u000b\u0001\u000b\u0005\u000b\u00c7\b\u000b\n\u000b\f\u000b"+
		"\u00ca\t\u000b\u0001\u000b\u0001\u000b\u0001\f\u0001\f\u0001\f\u0005\f"+
		"\u00d1\b\f\n\f\f\f\u00d4\t\f\u0001\f\u0001\f\u0001\r\u0001\r\u0001\r\u0001"+
		"\r\u0001\r\u0001\r\u0003\r\u00de\b\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001"+
		"\r\u0001\r\u0003\r\u00e6\b\r\u0001\u000e\u0001\u000e\u0001\u000e\u0001"+
		"\u000e\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0005\u000f\u00f0"+
		"\b\u000f\n\u000f\f\u000f\u00f3\t\u000f\u0001\u000f\u0003\u000f\u00f6\b"+
		"\u000f\u0001\u000f\u0001\u000f\u0001\u0010\u0001\u0010\u0001\u0010\u0001"+
		"\u0010\u0005\u0010\u00fe\b\u0010\n\u0010\f\u0010\u0101\t\u0010\u0001\u0011"+
		"\u0001\u0011\u0001\u0011\u0005\u0011\u0106\b\u0011\n\u0011\f\u0011\u0109"+
		"\t\u0011\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0005"+
		"\u0012\u0110\b\u0012\n\u0012\f\u0012\u0113\t\u0012\u0003\u0012\u0115\b"+
		"\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001"+
		"\u0012\u0005\u0012\u011d\b\u0012\n\u0012\f\u0012\u0120\t\u0012\u0003\u0012"+
		"\u0122\b\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012"+
		"\u0001\u0012\u0005\u0012\u012a\b\u0012\n\u0012\f\u0012\u012d\t\u0012\u0003"+
		"\u0012\u012f\b\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001"+
		"\u0012\u0001\u0012\u0005\u0012\u0137\b\u0012\n\u0012\f\u0012\u013a\t\u0012"+
		"\u0003\u0012\u013c\b\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012"+
		"\u0001\u0012\u0001\u0012\u0005\u0012\u0144\b\u0012\n\u0012\f\u0012\u0147"+
		"\t\u0012\u0003\u0012\u0149\b\u0012\u0001\u0012\u0003\u0012\u014c\b\u0012"+
		"\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0003\u0013\u0152\b\u0013"+
		"\u0001\u0013\u0001\u0013\u0003\u0013\u0156\b\u0013\u0001\u0013\u0001\u0013"+
		"\u0001\u0014\u0001\u0014\u0001\u0014\u0003\u0014\u015d\b\u0014\u0001\u0014"+
		"\u0001\u0014\u0001\u0015\u0001\u0015\u0001\u0015\u0005\u0015\u0164\b\u0015"+
		"\n\u0015\f\u0015\u0167\t\u0015\u0001\u0016\u0001\u0016\u0001\u0016\u0001"+
		"\u0017\u0001\u0017\u0001\u0017\u0005\u0017\u016f\b\u0017\n\u0017\f\u0017"+
		"\u0172\t\u0017\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018\u0001\u0018"+
		"\u0001\u0018\u0005\u0018\u017a\b\u0018\n\u0018\f\u0018\u017d\t\u0018\u0001"+
		"\u0018\u0001\u0018\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001"+
		"\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001"+
		"\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001"+
		"\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001"+
		"\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001"+
		"\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0003\u0019\u01a2"+
		"\b\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001"+
		"\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001"+
		"\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0005\u0019\u01b3\b\u0019\n"+
		"\u0019\f\u0019\u01b6\t\u0019\u0001\u001a\u0001\u001a\u0001\u001a\u0005"+
		"\u001a\u01bb\b\u001a\n\u001a\f\u001a\u01be\t\u001a\u0001\u001b\u0001\u001b"+
		"\u0001\u001c\u0001\u001c\u0001\u001c\u0001\u001c\u0001\u001c\u0003\u001c"+
		"\u01c7\b\u001c\u0001\u001d\u0001\u001d\u0001\u001d\u0005\u001d\u01cc\b"+
		"\u001d\n\u001d\f\u001d\u01cf\t\u001d\u0001\u001e\u0001\u001e\u0001\u001e"+
		"\u0001\u001e\u0003\u001e\u01d5\b\u001e\u0001\u001e\u0000\u00012\u001f"+
		"\u0000\u0002\u0004\u0006\b\n\f\u000e\u0010\u0012\u0014\u0016\u0018\u001a"+
		"\u001c\u001e \"$&(*,.02468:<\u0000\u0007\u0002\u0000))--\u0001\u0000&"+
		"\'\u0001\u000025\u0001\u000001\u0001\u0000./\u0001\u0000*,\u0001\u0000"+
		"()\u0207\u0000B\u0001\u0000\u0000\u0000\u0002G\u0001\u0000\u0000\u0000"+
		"\u0004M\u0001\u0000\u0000\u0000\u0006Y\u0001\u0000\u0000\u0000\b\u0082"+
		"\u0001\u0000\u0000\u0000\n\u0084\u0001\u0000\u0000\u0000\f\u0088\u0001"+
		"\u0000\u0000\u0000\u000e\u009e\u0001\u0000\u0000\u0000\u0010\u00a4\u0001"+
		"\u0000\u0000\u0000\u0012\u00ac\u0001\u0000\u0000\u0000\u0014\u00ae\u0001"+
		"\u0000\u0000\u0000\u0016\u00c1\u0001\u0000\u0000\u0000\u0018\u00cd\u0001"+
		"\u0000\u0000\u0000\u001a\u00e5\u0001\u0000\u0000\u0000\u001c\u00e7\u0001"+
		"\u0000\u0000\u0000\u001e\u00eb\u0001\u0000\u0000\u0000 \u00f9\u0001\u0000"+
		"\u0000\u0000\"\u0102\u0001\u0000\u0000\u0000$\u014b\u0001\u0000\u0000"+
		"\u0000&\u014d\u0001\u0000\u0000\u0000(\u0159\u0001\u0000\u0000\u0000*"+
		"\u0160\u0001\u0000\u0000\u0000,\u0168\u0001\u0000\u0000\u0000.\u016b\u0001"+
		"\u0000\u0000\u00000\u0173\u0001\u0000\u0000\u00002\u01a1\u0001\u0000\u0000"+
		"\u00004\u01b7\u0001\u0000\u0000\u00006\u01bf\u0001\u0000\u0000\u00008"+
		"\u01c6\u0001\u0000\u0000\u0000:\u01c8\u0001\u0000\u0000\u0000<\u01d4\u0001"+
		"\u0000\u0000\u0000>A\u0003&\u0013\u0000?A\u0003\u0002\u0001\u0000@>\u0001"+
		"\u0000\u0000\u0000@?\u0001\u0000\u0000\u0000AD\u0001\u0000\u0000\u0000"+
		"B@\u0001\u0000\u0000\u0000BC\u0001\u0000\u0000\u0000CE\u0001\u0000\u0000"+
		"\u0000DB\u0001\u0000\u0000\u0000EF\u0005\u0000\u0000\u0001F\u0001\u0001"+
		"\u0000\u0000\u0000GH\u0005\u0001\u0000\u0000HI\u0005\u0002\u0000\u0000"+
		"IJ\u00057\u0000\u0000JK\u00058\u0000\u0000KL\u0003\u0004\u0002\u0000L"+
		"\u0003\u0001\u0000\u0000\u0000MQ\u0005;\u0000\u0000NP\u0003\u0006\u0003"+
		"\u0000ON\u0001\u0000\u0000\u0000PS\u0001\u0000\u0000\u0000QO\u0001\u0000"+
		"\u0000\u0000QR\u0001\u0000\u0000\u0000RT\u0001\u0000\u0000\u0000SQ\u0001"+
		"\u0000\u0000\u0000TU\u0005<\u0000\u0000U\u0005\u0001\u0000\u0000\u0000"+
		"VZ\u0003\b\u0004\u0000WZ\u0003\u000e\u0007\u0000XZ\u0003&\u0013\u0000"+
		"YV\u0001\u0000\u0000\u0000YW\u0001\u0000\u0000\u0000YX\u0001\u0000\u0000"+
		"\u0000Z\u0007\u0001\u0000\u0000\u0000[\\\u0005\u0003\u0000\u0000\\^\u0005"+
		"#\u0000\u0000]_\u0005\b\u0000\u0000^]\u0001\u0000\u0000\u0000^_\u0001"+
		"\u0000\u0000\u0000_b\u0001\u0000\u0000\u0000`a\u00056\u0000\u0000ac\u0003"+
		"2\u0019\u0000b`\u0001\u0000\u0000\u0000bc\u0001\u0000\u0000\u0000c\u0083"+
		"\u0001\u0000\u0000\u0000de\u0005\u0003\u0000\u0000ef\u0005#\u0000\u0000"+
		"f\u0083\u0003\n\u0005\u0000gh\u0005#\u0000\u0000hi\u00056\u0000\u0000"+
		"ij\u0003\n\u0005\u0000jk\u0003\f\u0006\u0000k\u0083\u0001\u0000\u0000"+
		"\u0000lm\u0005#\u0000\u0000mn\u00056\u0000\u0000n\u0083\u0005#\u0000\u0000"+
		"or\u0005#\u0000\u0000pq\u00056\u0000\u0000qs\u00032\u0019\u0000rp\u0001"+
		"\u0000\u0000\u0000rs\u0001\u0000\u0000\u0000s\u0083\u0001\u0000\u0000"+
		"\u0000tu\u0005#\u0000\u0000uv\u00056\u0000\u0000vw\u0005\t\u0000\u0000"+
		"wx\u00057\u0000\u0000xy\u00032\u0019\u0000yz\u00058\u0000\u0000z\u0083"+
		"\u0001\u0000\u0000\u0000{|\u0005#\u0000\u0000|}\u00059\u0000\u0000}~\u0003"+
		"2\u0019\u0000~\u007f\u0005:\u0000\u0000\u007f\u0080\u00056\u0000\u0000"+
		"\u0080\u0081\u00032\u0019\u0000\u0081\u0083\u0001\u0000\u0000\u0000\u0082"+
		"[\u0001\u0000\u0000\u0000\u0082d\u0001\u0000\u0000\u0000\u0082g\u0001"+
		"\u0000\u0000\u0000\u0082l\u0001\u0000\u0000\u0000\u0082o\u0001\u0000\u0000"+
		"\u0000\u0082t\u0001\u0000\u0000\u0000\u0082{\u0001\u0000\u0000\u0000\u0083"+
		"\t\u0001\u0000\u0000\u0000\u0084\u0085\u00059\u0000\u0000\u0085\u0086"+
		"\u0005:\u0000\u0000\u0086\u0087\u0005\b\u0000\u0000\u0087\u000b\u0001"+
		"\u0000\u0000\u0000\u0088\u008a\u0005;\u0000\u0000\u0089\u008b\u0003:\u001d"+
		"\u0000\u008a\u0089\u0001\u0000\u0000\u0000\u008a\u008b\u0001\u0000\u0000"+
		"\u0000\u008b\u008c\u0001\u0000\u0000\u0000\u008c\u008d\u0005<\u0000\u0000"+
		"\u008d\r\u0001\u0000\u0000\u0000\u008e\u008f\u0005\"\u0000\u0000\u008f"+
		"\u0098\u00057\u0000\u0000\u0090\u0095\u00032\u0019\u0000\u0091\u0092\u0005"+
		"@\u0000\u0000\u0092\u0094\u00032\u0019\u0000\u0093\u0091\u0001\u0000\u0000"+
		"\u0000\u0094\u0097\u0001\u0000\u0000\u0000\u0095\u0093\u0001\u0000\u0000"+
		"\u0000\u0095\u0096\u0001\u0000\u0000\u0000\u0096\u0099\u0001\u0000\u0000"+
		"\u0000\u0097\u0095\u0001\u0000\u0000\u0000\u0098\u0090\u0001\u0000\u0000"+
		"\u0000\u0098\u0099\u0001\u0000\u0000\u0000\u0099\u009a\u0001\u0000\u0000"+
		"\u0000\u009a\u009f\u00058\u0000\u0000\u009b\u009f\u00032\u0019\u0000\u009c"+
		"\u009f\u0003\u0010\b\u0000\u009d\u009f\u0003\u0012\t\u0000\u009e\u008e"+
		"\u0001\u0000\u0000\u0000\u009e\u009b\u0001\u0000\u0000\u0000\u009e\u009c"+
		"\u0001\u0000\u0000\u0000\u009e\u009d\u0001\u0000\u0000\u0000\u009f\u000f"+
		"\u0001\u0000\u0000\u0000\u00a0\u00a5\u0003\u0014\n\u0000\u00a1\u00a5\u0003"+
		"\u001a\r\u0000\u00a2\u00a5\u0003\u001e\u000f\u0000\u00a3\u00a5\u00030"+
		"\u0018\u0000\u00a4\u00a0\u0001\u0000\u0000\u0000\u00a4\u00a1\u0001\u0000"+
		"\u0000\u0000\u00a4\u00a2\u0001\u0000\u0000\u0000\u00a4\u00a3\u0001\u0000"+
		"\u0000\u0000\u00a5\u0011\u0001\u0000\u0000\u0000\u00a6\u00ad\u0005\u0015"+
		"\u0000\u0000\u00a7\u00ad\u0005\u0016\u0000\u0000\u00a8\u00aa\u0005\u0017"+
		"\u0000\u0000\u00a9\u00ab\u00032\u0019\u0000\u00aa\u00a9\u0001\u0000\u0000"+
		"\u0000\u00aa\u00ab\u0001\u0000\u0000\u0000\u00ab\u00ad\u0001\u0000\u0000"+
		"\u0000\u00ac\u00a6\u0001\u0000\u0000\u0000\u00ac\u00a7\u0001\u0000\u0000"+
		"\u0000\u00ac\u00a8\u0001\u0000\u0000\u0000\u00ad\u0013\u0001\u0000\u0000"+
		"\u0000\u00ae\u00af\u0005\u000f\u0000\u0000\u00af\u00b0\u00032\u0019\u0000"+
		"\u00b0\u00b4\u0005;\u0000\u0000\u00b1\u00b3\u0003\u0006\u0003\u0000\u00b2"+
		"\u00b1\u0001\u0000\u0000\u0000\u00b3\u00b6\u0001\u0000\u0000\u0000\u00b4"+
		"\u00b2\u0001\u0000\u0000\u0000\u00b4\u00b5\u0001\u0000\u0000\u0000\u00b5"+
		"\u00b7\u0001\u0000\u0000\u0000\u00b6\u00b4\u0001\u0000\u0000\u0000\u00b7"+
		"\u00bb\u0005<\u0000\u0000\u00b8\u00ba\u0003\u0016\u000b\u0000\u00b9\u00b8"+
		"\u0001\u0000\u0000\u0000\u00ba\u00bd\u0001\u0000\u0000\u0000\u00bb\u00b9"+
		"\u0001\u0000\u0000\u0000\u00bb\u00bc\u0001\u0000\u0000\u0000\u00bc\u00bf"+
		"\u0001\u0000\u0000\u0000\u00bd\u00bb\u0001\u0000\u0000\u0000\u00be\u00c0"+
		"\u0003\u0018\f\u0000\u00bf\u00be\u0001\u0000\u0000\u0000\u00bf\u00c0\u0001"+
		"\u0000\u0000\u0000\u00c0\u0015\u0001\u0000\u0000\u0000\u00c1\u00c2\u0005"+
		"\u0010\u0000\u0000\u00c2\u00c3\u0005\u000f\u0000\u0000\u00c3\u00c4\u0003"+
		"2\u0019\u0000\u00c4\u00c8\u0005;\u0000\u0000\u00c5\u00c7\u0003\u0006\u0003"+
		"\u0000\u00c6\u00c5\u0001\u0000\u0000\u0000\u00c7\u00ca\u0001\u0000\u0000"+
		"\u0000\u00c8\u00c6\u0001\u0000\u0000\u0000\u00c8\u00c9\u0001\u0000\u0000"+
		"\u0000\u00c9\u00cb\u0001\u0000\u0000\u0000\u00ca\u00c8\u0001\u0000\u0000"+
		"\u0000\u00cb\u00cc\u0005<\u0000\u0000\u00cc\u0017\u0001\u0000\u0000\u0000"+
		"\u00cd\u00ce\u0005\u0010\u0000\u0000\u00ce\u00d2\u0005;\u0000\u0000\u00cf"+
		"\u00d1\u0003\u0006\u0003\u0000\u00d0\u00cf\u0001\u0000\u0000\u0000\u00d1"+
		"\u00d4\u0001\u0000\u0000\u0000\u00d2\u00d0\u0001\u0000\u0000\u0000\u00d2"+
		"\u00d3\u0001\u0000\u0000\u0000\u00d3\u00d5\u0001\u0000\u0000\u0000\u00d4"+
		"\u00d2\u0001\u0000\u0000\u0000\u00d5\u00d6\u0005<\u0000\u0000\u00d6\u0019"+
		"\u0001\u0000\u0000\u0000\u00d7\u00d8\u0005\u0011\u0000\u0000\u00d8\u00d9"+
		"\u0003\u001c\u000e\u0000\u00d9\u00da\u0005=\u0000\u0000\u00da\u00db\u0003"+
		"2\u0019\u0000\u00db\u00dd\u0005=\u0000\u0000\u00dc\u00de\u0003\u000e\u0007"+
		"\u0000\u00dd\u00dc\u0001\u0000\u0000\u0000\u00dd\u00de\u0001\u0000\u0000"+
		"\u0000\u00de\u00df\u0001\u0000\u0000\u0000\u00df\u00e0\u0003\u0004\u0002"+
		"\u0000\u00e0\u00e6\u0001\u0000\u0000\u0000\u00e1\u00e2\u0005\u0011\u0000"+
		"\u0000\u00e2\u00e3\u00032\u0019\u0000\u00e3\u00e4\u0003\u0004\u0002\u0000"+
		"\u00e4\u00e6\u0001\u0000\u0000\u0000\u00e5\u00d7\u0001\u0000\u0000\u0000"+
		"\u00e5\u00e1\u0001\u0000\u0000\u0000\u00e6\u001b\u0001\u0000\u0000\u0000"+
		"\u00e7\u00e8\u0005#\u0000\u0000\u00e8\u00e9\u00056\u0000\u0000\u00e9\u00ea"+
		"\u00032\u0019\u0000\u00ea\u001d\u0001\u0000\u0000\u0000\u00eb\u00ec\u0005"+
		"\u0012\u0000\u0000\u00ec\u00ed\u00032\u0019\u0000\u00ed\u00f1\u0005;\u0000"+
		"\u0000\u00ee\u00f0\u0003 \u0010\u0000\u00ef\u00ee\u0001\u0000\u0000\u0000"+
		"\u00f0\u00f3\u0001\u0000\u0000\u0000\u00f1\u00ef\u0001\u0000\u0000\u0000"+
		"\u00f1\u00f2\u0001\u0000\u0000\u0000\u00f2\u00f5\u0001\u0000\u0000\u0000"+
		"\u00f3\u00f1\u0001\u0000\u0000\u0000\u00f4\u00f6\u0003\"\u0011\u0000\u00f5"+
		"\u00f4\u0001\u0000\u0000\u0000\u00f5\u00f6\u0001\u0000\u0000\u0000\u00f6"+
		"\u00f7\u0001\u0000\u0000\u0000\u00f7\u00f8\u0005<\u0000\u0000\u00f8\u001f"+
		"\u0001\u0000\u0000\u0000\u00f9\u00fa\u0005\u0004\u0000\u0000\u00fa\u00fb"+
		"\u00032\u0019\u0000\u00fb\u00ff\u0005>\u0000\u0000\u00fc\u00fe\u0003\u0006"+
		"\u0003\u0000\u00fd\u00fc\u0001\u0000\u0000\u0000\u00fe\u0101\u0001\u0000"+
		"\u0000\u0000\u00ff\u00fd\u0001\u0000\u0000\u0000\u00ff\u0100\u0001\u0000"+
		"\u0000\u0000\u0100!\u0001\u0000\u0000\u0000\u0101\u00ff\u0001\u0000\u0000"+
		"\u0000\u0102\u0103\u0005\u0005\u0000\u0000\u0103\u0107\u0005>\u0000\u0000"+
		"\u0104\u0106\u0003\u0006\u0003\u0000\u0105\u0104\u0001\u0000\u0000\u0000"+
		"\u0106\u0109\u0001\u0000\u0000\u0000\u0107\u0105\u0001\u0000\u0000\u0000"+
		"\u0107\u0108\u0001\u0000\u0000\u0000\u0108#\u0001\u0000\u0000\u0000\u0109"+
		"\u0107\u0001\u0000\u0000\u0000\u010a\u010b\u0005\u0013\u0000\u0000\u010b"+
		"\u0114\u00057\u0000\u0000\u010c\u0111\u00032\u0019\u0000\u010d\u010e\u0005"+
		"@\u0000\u0000\u010e\u0110\u00032\u0019\u0000\u010f\u010d\u0001\u0000\u0000"+
		"\u0000\u0110\u0113\u0001\u0000\u0000\u0000\u0111\u010f\u0001\u0000\u0000"+
		"\u0000\u0111\u0112\u0001\u0000\u0000\u0000\u0112\u0115\u0001\u0000\u0000"+
		"\u0000\u0113\u0111\u0001\u0000\u0000\u0000\u0114\u010c\u0001\u0000\u0000"+
		"\u0000\u0114\u0115\u0001\u0000\u0000\u0000\u0115\u0116\u0001\u0000\u0000"+
		"\u0000\u0116\u014c\u00058\u0000\u0000\u0117\u0118\u0005\u0014\u0000\u0000"+
		"\u0118\u0121\u00057\u0000\u0000\u0119\u011e\u00032\u0019\u0000\u011a\u011b"+
		"\u0005@\u0000\u0000\u011b\u011d\u00032\u0019\u0000\u011c\u011a\u0001\u0000"+
		"\u0000\u0000\u011d\u0120\u0001\u0000\u0000\u0000\u011e\u011c\u0001\u0000"+
		"\u0000\u0000\u011e\u011f\u0001\u0000\u0000\u0000\u011f\u0122\u0001\u0000"+
		"\u0000\u0000\u0120\u011e\u0001\u0000\u0000\u0000\u0121\u0119\u0001\u0000"+
		"\u0000\u0000\u0121\u0122\u0001\u0000\u0000\u0000\u0122\u0123\u0001\u0000"+
		"\u0000\u0000\u0123\u014c\u00058\u0000\u0000\u0124\u0125\u0005#\u0000\u0000"+
		"\u0125\u012e\u00057\u0000\u0000\u0126\u012b\u00032\u0019\u0000\u0127\u0128"+
		"\u0005@\u0000\u0000\u0128\u012a\u00032\u0019\u0000\u0129\u0127\u0001\u0000"+
		"\u0000\u0000\u012a\u012d\u0001\u0000\u0000\u0000\u012b\u0129\u0001\u0000"+
		"\u0000\u0000\u012b\u012c\u0001\u0000\u0000\u0000\u012c\u012f\u0001\u0000"+
		"\u0000\u0000\u012d\u012b\u0001\u0000\u0000\u0000\u012e\u0126\u0001\u0000"+
		"\u0000\u0000\u012e\u012f\u0001\u0000\u0000\u0000\u012f\u0130\u0001\u0000"+
		"\u0000\u0000\u0130\u014c\u00058\u0000\u0000\u0131\u0132\u0005\f\u0000"+
		"\u0000\u0132\u013b\u00057\u0000\u0000\u0133\u0138\u00032\u0019\u0000\u0134"+
		"\u0135\u0005@\u0000\u0000\u0135\u0137\u00032\u0019\u0000\u0136\u0134\u0001"+
		"\u0000\u0000\u0000\u0137\u013a\u0001\u0000\u0000\u0000\u0138\u0136\u0001"+
		"\u0000\u0000\u0000\u0138\u0139\u0001\u0000\u0000\u0000\u0139\u013c\u0001"+
		"\u0000\u0000\u0000\u013a\u0138\u0001\u0000\u0000\u0000\u013b\u0133\u0001"+
		"\u0000\u0000\u0000\u013b\u013c\u0001\u0000\u0000\u0000\u013c\u013d\u0001"+
		"\u0000\u0000\u0000\u013d\u014c\u00058\u0000\u0000\u013e\u013f\u0005\u000e"+
		"\u0000\u0000\u013f\u0148\u00057\u0000\u0000\u0140\u0145\u00032\u0019\u0000"+
		"\u0141\u0142\u0005@\u0000\u0000\u0142\u0144\u00032\u0019\u0000\u0143\u0141"+
		"\u0001\u0000\u0000\u0000\u0144\u0147\u0001\u0000\u0000\u0000\u0145\u0143"+
		"\u0001\u0000\u0000\u0000\u0145\u0146\u0001\u0000\u0000\u0000\u0146\u0149"+
		"\u0001\u0000\u0000\u0000\u0147\u0145\u0001\u0000\u0000\u0000\u0148\u0140"+
		"\u0001\u0000\u0000\u0000\u0148\u0149\u0001\u0000\u0000\u0000\u0149\u014a"+
		"\u0001\u0000\u0000\u0000\u014a\u014c\u00058\u0000\u0000\u014b\u010a\u0001"+
		"\u0000\u0000\u0000\u014b\u0117\u0001\u0000\u0000\u0000\u014b\u0124\u0001"+
		"\u0000\u0000\u0000\u014b\u0131\u0001\u0000\u0000\u0000\u014b\u013e\u0001"+
		"\u0000\u0000\u0000\u014c%\u0001\u0000\u0000\u0000\u014d\u014e\u0005\u0006"+
		"\u0000\u0000\u014e\u014f\u0005#\u0000\u0000\u014f\u0151\u00057\u0000\u0000"+
		"\u0150\u0152\u0003*\u0015\u0000\u0151\u0150\u0001\u0000\u0000\u0000\u0151"+
		"\u0152\u0001\u0000\u0000\u0000\u0152\u0153\u0001\u0000\u0000\u0000\u0153"+
		"\u0155\u00058\u0000\u0000\u0154\u0156\u0005\b\u0000\u0000\u0155\u0154"+
		"\u0001\u0000\u0000\u0000\u0155\u0156\u0001\u0000\u0000\u0000\u0156\u0157"+
		"\u0001\u0000\u0000\u0000\u0157\u0158\u0003\u0004\u0002\u0000\u0158\'\u0001"+
		"\u0000\u0000\u0000\u0159\u015a\u0005#\u0000\u0000\u015a\u015c\u00057\u0000"+
		"\u0000\u015b\u015d\u0003.\u0017\u0000\u015c\u015b\u0001\u0000\u0000\u0000"+
		"\u015c\u015d\u0001\u0000\u0000\u0000\u015d\u015e\u0001\u0000\u0000\u0000"+
		"\u015e\u015f\u00058\u0000\u0000\u015f)\u0001\u0000\u0000\u0000\u0160\u0165"+
		"\u0003,\u0016\u0000\u0161\u0162\u0005@\u0000\u0000\u0162\u0164\u0003,"+
		"\u0016\u0000\u0163\u0161\u0001\u0000\u0000\u0000\u0164\u0167\u0001\u0000"+
		"\u0000\u0000\u0165\u0163\u0001\u0000\u0000\u0000\u0165\u0166\u0001\u0000"+
		"\u0000\u0000\u0166+\u0001\u0000\u0000\u0000\u0167\u0165\u0001\u0000\u0000"+
		"\u0000\u0168\u0169\u0005#\u0000\u0000\u0169\u016a\u0005\b\u0000\u0000"+
		"\u016a-\u0001\u0000\u0000\u0000\u016b\u0170\u00032\u0019\u0000\u016c\u016d"+
		"\u0005@\u0000\u0000\u016d\u016f\u00032\u0019\u0000\u016e\u016c\u0001\u0000"+
		"\u0000\u0000\u016f\u0172\u0001\u0000\u0000\u0000\u0170\u016e\u0001\u0000"+
		"\u0000\u0000\u0170\u0171\u0001\u0000\u0000\u0000\u0171/\u0001\u0000\u0000"+
		"\u0000\u0172\u0170\u0001\u0000\u0000\u0000\u0173\u0174\u0005\u0007\u0000"+
		"\u0000\u0174\u0175\u00057\u0000\u0000\u0175\u0176\u00032\u0019\u0000\u0176"+
		"\u0177\u00058\u0000\u0000\u0177\u017b\u00059\u0000\u0000\u0178\u017a\u0003"+
		"\u0006\u0003\u0000\u0179\u0178\u0001\u0000\u0000\u0000\u017a\u017d\u0001"+
		"\u0000\u0000\u0000\u017b\u0179\u0001\u0000\u0000\u0000\u017b\u017c\u0001"+
		"\u0000\u0000\u0000\u017c\u017e\u0001\u0000\u0000\u0000\u017d\u017b\u0001"+
		"\u0000\u0000\u0000\u017e\u017f\u0005:\u0000\u0000\u017f1\u0001\u0000\u0000"+
		"\u0000\u0180\u0181\u0006\u0019\uffff\uffff\u0000\u0181\u0182\u0007\u0000"+
		"\u0000\u0000\u0182\u01a2\u00032\u0019\f\u0183\u01a2\u00038\u001c\u0000"+
		"\u0184\u0185\u00057\u0000\u0000\u0185\u0186\u00032\u0019\u0000\u0186\u0187"+
		"\u00058\u0000\u0000\u0187\u01a2\u0001\u0000\u0000\u0000\u0188\u0189\u0005"+
		"9\u0000\u0000\u0189\u018a\u00032\u0019\u0000\u018a\u018b\u0005:\u0000"+
		"\u0000\u018b\u01a2\u0001\u0000\u0000\u0000\u018c\u018d\u0005#\u0000\u0000"+
		"\u018d\u018e\u00059\u0000\u0000\u018e\u018f\u00032\u0019\u0000\u018f\u0190"+
		"\u0005:\u0000\u0000\u0190\u01a2\u0001\u0000\u0000\u0000\u0191\u01a2\u0003"+
		"$\u0012\u0000\u0192\u01a2\u0005#\u0000\u0000\u0193\u01a2\u0003<\u001e"+
		"\u0000\u0194\u0195\u0005#\u0000\u0000\u0195\u0196\u0005?\u0000\u0000\u0196"+
		"\u01a2\u0005#\u0000\u0000\u0197\u0198\u0005#\u0000\u0000\u0198\u0199\u0005"+
		"?\u0000\u0000\u0199\u01a2\u00032\u0019\u0003\u019a\u019b\u0005#\u0000"+
		"\u0000\u019b\u019c\u0005\b\u0000\u0000\u019c\u019d\u00056\u0000\u0000"+
		"\u019d\u01a2\u00032\u0019\u0002\u019e\u019f\u0005#\u0000\u0000\u019f\u01a0"+
		"\u0007\u0001\u0000\u0000\u01a0\u01a2\u00032\u0019\u0001\u01a1\u0180\u0001"+
		"\u0000\u0000\u0000\u01a1\u0183\u0001\u0000\u0000\u0000\u01a1\u0184\u0001"+
		"\u0000\u0000\u0000\u01a1\u0188\u0001\u0000\u0000\u0000\u01a1\u018c\u0001"+
		"\u0000\u0000\u0000\u01a1\u0191\u0001\u0000\u0000\u0000\u01a1\u0192\u0001"+
		"\u0000\u0000\u0000\u01a1\u0193\u0001\u0000\u0000\u0000\u01a1\u0194\u0001"+
		"\u0000\u0000\u0000\u01a1\u0197\u0001\u0000\u0000\u0000\u01a1\u019a\u0001"+
		"\u0000\u0000\u0000\u01a1\u019e\u0001\u0000\u0000\u0000\u01a2\u01b4\u0001"+
		"\u0000\u0000\u0000\u01a3\u01a4\n\u0011\u0000\u0000\u01a4\u01a5\u0007\u0002"+
		"\u0000\u0000\u01a5\u01b3\u00032\u0019\u0012\u01a6\u01a7\n\u0010\u0000"+
		"\u0000\u01a7\u01a8\u0007\u0003\u0000\u0000\u01a8\u01b3\u00032\u0019\u0011"+
		"\u01a9\u01aa\n\u000f\u0000\u0000\u01aa\u01ab\u0007\u0004\u0000\u0000\u01ab"+
		"\u01b3\u00032\u0019\u0010\u01ac\u01ad\n\u000e\u0000\u0000\u01ad\u01ae"+
		"\u0007\u0005\u0000\u0000\u01ae\u01b3\u00032\u0019\u000f\u01af\u01b0\n"+
		"\r\u0000\u0000\u01b0\u01b1\u0007\u0006\u0000\u0000\u01b1\u01b3\u00032"+
		"\u0019\u000e\u01b2\u01a3\u0001\u0000\u0000\u0000\u01b2\u01a6\u0001\u0000"+
		"\u0000\u0000\u01b2\u01a9\u0001\u0000\u0000\u0000\u01b2\u01ac\u0001\u0000"+
		"\u0000\u0000\u01b2\u01af\u0001\u0000\u0000\u0000\u01b3\u01b6\u0001\u0000"+
		"\u0000\u0000\u01b4\u01b2\u0001\u0000\u0000\u0000\u01b4\u01b5\u0001\u0000"+
		"\u0000\u0000\u01b53\u0001\u0000\u0000\u0000\u01b6\u01b4\u0001\u0000\u0000"+
		"\u0000\u01b7\u01bc\u00032\u0019\u0000\u01b8\u01b9\u0005@\u0000\u0000\u01b9"+
		"\u01bb\u00032\u0019\u0000\u01ba\u01b8\u0001\u0000\u0000\u0000\u01bb\u01be"+
		"\u0001\u0000\u0000\u0000\u01bc\u01ba\u0001\u0000\u0000\u0000\u01bc\u01bd"+
		"\u0001\u0000\u0000\u0000\u01bd5\u0001\u0000\u0000\u0000\u01be\u01bc\u0001"+
		"\u0000\u0000\u0000\u01bf\u01c0\u00038\u001c\u0000\u01c07\u0001\u0000\u0000"+
		"\u0000\u01c1\u01c7\u0005\u0019\u0000\u0000\u01c2\u01c7\u0005\u001a\u0000"+
		"\u0000\u01c3\u01c7\u0005\u001b\u0000\u0000\u01c4\u01c7\u0005\u0018\u0000"+
		"\u0000\u01c5\u01c7\u0005\u001c\u0000\u0000\u01c6\u01c1\u0001\u0000\u0000"+
		"\u0000\u01c6\u01c2\u0001\u0000\u0000\u0000\u01c6\u01c3\u0001\u0000\u0000"+
		"\u0000\u01c6\u01c4\u0001\u0000\u0000\u0000\u01c6\u01c5\u0001\u0000\u0000"+
		"\u0000\u01c79\u0001\u0000\u0000\u0000\u01c8\u01cd\u00032\u0019\u0000\u01c9"+
		"\u01ca\u0005@\u0000\u0000\u01ca\u01cc\u00032\u0019\u0000\u01cb\u01c9\u0001"+
		"\u0000\u0000\u0000\u01cc\u01cf\u0001\u0000\u0000\u0000\u01cd\u01cb\u0001"+
		"\u0000\u0000\u0000\u01cd\u01ce\u0001\u0000\u0000\u0000\u01ce;\u0001\u0000"+
		"\u0000\u0000\u01cf\u01cd\u0001\u0000\u0000\u0000\u01d0\u01d1\u0005#\u0000"+
		"\u0000\u01d1\u01d5\u0005$\u0000\u0000\u01d2\u01d3\u0005#\u0000\u0000\u01d3"+
		"\u01d5\u0005%\u0000\u0000\u01d4\u01d0\u0001\u0000\u0000\u0000\u01d4\u01d2"+
		"\u0001\u0000\u0000\u0000\u01d5=\u0001\u0000\u0000\u00002@BQY^br\u0082"+
		"\u008a\u0095\u0098\u009e\u00a4\u00aa\u00ac\u00b4\u00bb\u00bf\u00c8\u00d2"+
		"\u00dd\u00e5\u00f1\u00f5\u00ff\u0107\u0111\u0114\u011e\u0121\u012b\u012e"+
		"\u0138\u013b\u0145\u0148\u014b\u0151\u0155\u015c\u0165\u0170\u017b\u01a1"+
		"\u01b2\u01b4\u01bc\u01c6\u01cd\u01d4";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}