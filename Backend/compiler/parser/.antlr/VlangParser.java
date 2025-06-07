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
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, TIPO=7, CONDICIONAL_ETC=8, 
		LEN=9, CAP=10, APPEND=11, IF=12, ELSE=13, FOR=14, SWITCH=15, INDEXOF=16, 
		JOIN=17, BOOLEANO=18, ENTERO=19, DECIMAL=20, CADENA=21, CARACTER=22, TIPO_ENTERO=23, 
		TIPO_DECIMAL=24, TIPO_CADENA=25, TIPO_BOOLEANO=26, TIPO_CHAR=27, PRINT=28, 
		ID=29, INC=30, DEC=31, SUMAIMPLICITA=32, RESTOIMPLICITO=33, PLUS=34, MINUS=35, 
		MUL=36, DIV=37, MOD=38, NOT=39, OR=40, AND=41, EQ=42, NEQ=43, LE=44, GE=45, 
		LT=46, GT=47, ASSIGN=48, LPAREN=49, RPAREN=50, LBRACK=51, RBRACK=52, LBRACE=53, 
		RBRACE=54, SEMICOLON=55, COLON=56, DOT=57, COMMA=58, WS=59, LINE_COMMENT=60, 
		BLOCK_COMMENT=61;
	public static final int
		RULE_programa = 0, RULE_funcMain = 1, RULE_block = 2, RULE_declaraciones = 3, 
		RULE_varDcl = 4, RULE_tipoSlice = 5, RULE_stmt = 6, RULE_sentencias_control = 7, 
		RULE_ifDcl = 8, RULE_forDcl = 9, RULE_switchDcl = 10, RULE_caseBlock = 11, 
		RULE_defaultBlock = 12, RULE_llamadaFuncion = 13, RULE_whileDcl = 14, 
		RULE_expresion = 15, RULE_parametros = 16, RULE_valores = 17, RULE_valor = 18, 
		RULE_listaExpresiones = 19, RULE_incredecre = 20;
	private static String[] makeRuleNames() {
		return new String[] {
			"programa", "funcMain", "block", "declaraciones", "varDcl", "tipoSlice", 
			"stmt", "sentencias_control", "ifDcl", "forDcl", "switchDcl", "caseBlock", 
			"defaultBlock", "llamadaFuncion", "whileDcl", "expresion", "parametros", 
			"valores", "valor", "listaExpresiones", "incredecre"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'fn'", "'main'", "'mut'", "'case'", "'default'", "'while'", null, 
			null, "'len'", "'cap'", "'append'", "'if'", "'else'", "'for'", "'switch'", 
			"'indexOf'", "'join'", null, null, null, null, null, "'int'", "'float64'", 
			"'string'", "'bool'", "'rune'", "'print'", null, "'++'", "'--'", "'+='", 
			"'-='", "'+'", "'-'", "'*'", "'/'", "'%'", "'!'", "'||'", "'&&'", "'=='", 
			"'!='", "'<='", "'>='", "'<'", "'>'", "'='", "'('", "')'", "'['", "']'", 
			"'{'", "'}'", "';'", "':'", "'.'", "','"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, "TIPO", "CONDICIONAL_ETC", 
			"LEN", "CAP", "APPEND", "IF", "ELSE", "FOR", "SWITCH", "INDEXOF", "JOIN", 
			"BOOLEANO", "ENTERO", "DECIMAL", "CADENA", "CARACTER", "TIPO_ENTERO", 
			"TIPO_DECIMAL", "TIPO_CADENA", "TIPO_BOOLEANO", "TIPO_CHAR", "PRINT", 
			"ID", "INC", "DEC", "SUMAIMPLICITA", "RESTOIMPLICITO", "PLUS", "MINUS", 
			"MUL", "DIV", "MOD", "NOT", "OR", "AND", "EQ", "NEQ", "LE", "GE", "LT", 
			"GT", "ASSIGN", "LPAREN", "RPAREN", "LBRACK", "RBRACK", "LBRACE", "RBRACE", 
			"SEMICOLON", "COLON", "DOT", "COMMA", "WS", "LINE_COMMENT", "BLOCK_COMMENT"
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
			setState(45);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__0) {
				{
				{
				setState(42);
				funcMain();
				}
				}
				setState(47);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(48);
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
			setState(50);
			match(T__0);
			setState(51);
			match(T__1);
			setState(52);
			match(LPAREN);
			setState(53);
			match(RPAREN);
			setState(54);
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
			setState(56);
			match(LBRACE);
			setState(60);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 2815334696288584L) != 0)) {
				{
				{
				setState(57);
				declaraciones();
				}
				}
				setState(62);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(63);
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
		public DeclaracionesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaraciones; }
	}

	public final DeclaracionesContext declaraciones() throws RecognitionException {
		DeclaracionesContext _localctx = new DeclaracionesContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_declaraciones);
		try {
			setState(67);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__2:
				enterOuterAlt(_localctx, 1);
				{
				setState(65);
				varDcl();
				}
				break;
			case T__5:
			case CONDICIONAL_ETC:
			case INDEXOF:
			case JOIN:
			case BOOLEANO:
			case ENTERO:
			case DECIMAL:
			case CADENA:
			case CARACTER:
			case PRINT:
			case ID:
			case MINUS:
			case NOT:
			case LPAREN:
			case LBRACK:
				enterOuterAlt(_localctx, 2);
				{
				setState(66);
				stmt();
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
	public static class VariableDeclarationContext extends VarDclContext {
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public TerminalNode TIPO() { return getToken(VlangParser.TIPO, 0); }
		public TerminalNode ASSIGN() { return getToken(VlangParser.ASSIGN, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public VariableDeclarationContext(VarDclContext ctx) { copyFrom(ctx); }
	}

	public final VarDclContext varDcl() throws RecognitionException {
		VarDclContext _localctx = new VarDclContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_varDcl);
		int _la;
		try {
			_localctx = new VariableDeclarationContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(69);
			match(T__2);
			setState(70);
			match(ID);
			setState(72);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==TIPO) {
				{
				setState(71);
				match(TIPO);
				}
			}

			setState(76);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ASSIGN) {
				{
				setState(74);
				match(ASSIGN);
				setState(75);
				expresion(0);
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
	public static class TipoSliceContext extends ParserRuleContext {
		public TerminalNode LBRACK() { return getToken(VlangParser.LBRACK, 0); }
		public TerminalNode RBRACK() { return getToken(VlangParser.RBRACK, 0); }
		public TerminalNode TIPO() { return getToken(VlangParser.TIPO, 0); }
		public TipoSliceContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tipoSlice; }
	}

	public final TipoSliceContext tipoSlice() throws RecognitionException {
		TipoSliceContext _localctx = new TipoSliceContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_tipoSlice);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(78);
			match(LBRACK);
			setState(79);
			match(RBRACK);
			setState(80);
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

	public final StmtContext stmt() throws RecognitionException {
		StmtContext _localctx = new StmtContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_stmt);
		int _la;
		try {
			setState(97);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case PRINT:
				_localctx = new PrintStatementContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(82);
				match(PRINT);
				setState(83);
				match(LPAREN);
				setState(92);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 2815334427852800L) != 0)) {
					{
					setState(84);
					expresion(0);
					setState(89);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==COMMA) {
						{
						{
						setState(85);
						match(COMMA);
						setState(86);
						expresion(0);
						}
						}
						setState(91);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					}
				}

				setState(94);
				match(RPAREN);
				}
				break;
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
				setState(95);
				expresion(0);
				}
				break;
			case T__5:
			case CONDICIONAL_ETC:
				_localctx = new ControlStatementContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(96);
				sentencias_control();
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
		enterRule(_localctx, 14, RULE_sentencias_control);
		try {
			setState(103);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				_localctx = new If_contextContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(99);
				ifDcl();
				}
				break;
			case 2:
				_localctx = new For_contextContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(100);
				forDcl();
				}
				break;
			case 3:
				_localctx = new Switch_contextContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(101);
				switchDcl();
				}
				break;
			case 4:
				_localctx = new While_contextContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(102);
				whileDcl();
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
	public static class IfDclContext extends ParserRuleContext {
		public List<TerminalNode> CONDICIONAL_ETC() { return getTokens(VlangParser.CONDICIONAL_ETC); }
		public TerminalNode CONDICIONAL_ETC(int i) {
			return getToken(VlangParser.CONDICIONAL_ETC, i);
		}
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public List<TerminalNode> LBRACE() { return getTokens(VlangParser.LBRACE); }
		public TerminalNode LBRACE(int i) {
			return getToken(VlangParser.LBRACE, i);
		}
		public List<TerminalNode> RBRACE() { return getTokens(VlangParser.RBRACE); }
		public TerminalNode RBRACE(int i) {
			return getToken(VlangParser.RBRACE, i);
		}
		public TerminalNode LPAREN() { return getToken(VlangParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(VlangParser.RPAREN, 0); }
		public List<DeclaracionesContext> declaraciones() {
			return getRuleContexts(DeclaracionesContext.class);
		}
		public DeclaracionesContext declaraciones(int i) {
			return getRuleContext(DeclaracionesContext.class,i);
		}
		public IfDclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ifDcl; }
	}

	public final IfDclContext ifDcl() throws RecognitionException {
		IfDclContext _localctx = new IfDclContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_ifDcl);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(105);
			match(CONDICIONAL_ETC);
			setState(107);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
			case 1:
				{
				setState(106);
				match(LPAREN);
				}
				break;
			}
			setState(109);
			expresion(0);
			setState(111);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==RPAREN) {
				{
				setState(110);
				match(RPAREN);
				}
			}

			setState(113);
			match(LBRACE);
			setState(117);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 2815334696288584L) != 0)) {
				{
				{
				setState(114);
				declaraciones();
				}
				}
				setState(119);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(120);
			match(RBRACE);
			setState(130);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
			case 1:
				{
				setState(121);
				match(CONDICIONAL_ETC);
				setState(122);
				match(LBRACE);
				setState(126);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 2815334696288584L) != 0)) {
					{
					{
					setState(123);
					declaraciones();
					}
					}
					setState(128);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(129);
				match(RBRACE);
				}
				break;
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
		public TerminalNode CONDICIONAL_ETC() { return getToken(VlangParser.CONDICIONAL_ETC, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public BlockContext block() {
			return getRuleContext(BlockContext.class,0);
		}
		public TerminalNode LPAREN() { return getToken(VlangParser.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(VlangParser.RPAREN, 0); }
		public ForCondicionUnicaContext(ForDclContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ForClasicoContext extends ForDclContext {
		public TerminalNode CONDICIONAL_ETC() { return getToken(VlangParser.CONDICIONAL_ETC, 0); }
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
		public TerminalNode LPAREN() { return getToken(VlangParser.LPAREN, 0); }
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public ForClasicoContext(ForDclContext ctx) { copyFrom(ctx); }
	}

	public final ForDclContext forDcl() throws RecognitionException {
		ForDclContext _localctx = new ForDclContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_forDcl);
		int _la;
		try {
			setState(157);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,19,_ctx) ) {
			case 1:
				_localctx = new ForClasicoContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(132);
				match(CONDICIONAL_ETC);
				setState(134);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,14,_ctx) ) {
				case 1:
					{
					setState(133);
					match(LPAREN);
					}
					break;
				}
				setState(137);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 2815334696288576L) != 0)) {
					{
					setState(136);
					stmt();
					}
				}

				setState(139);
				match(SEMICOLON);
				setState(140);
				expresion(0);
				setState(141);
				match(SEMICOLON);
				setState(143);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 2815334696288576L) != 0)) {
					{
					setState(142);
					stmt();
					}
				}

				setState(145);
				block();
				}
				break;
			case 2:
				_localctx = new ForCondicionUnicaContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(147);
				match(CONDICIONAL_ETC);
				setState(149);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,17,_ctx) ) {
				case 1:
					{
					setState(148);
					match(LPAREN);
					}
					break;
				}
				setState(151);
				expresion(0);
				setState(153);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==RPAREN) {
					{
					setState(152);
					match(RPAREN);
					}
				}

				setState(155);
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
	public static class SwitchDclContext extends ParserRuleContext {
		public TerminalNode CONDICIONAL_ETC() { return getToken(VlangParser.CONDICIONAL_ETC, 0); }
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
		enterRule(_localctx, 20, RULE_switchDcl);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(159);
			match(CONDICIONAL_ETC);
			setState(160);
			expresion(0);
			setState(161);
			match(LBRACE);
			setState(165);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__3) {
				{
				{
				setState(162);
				caseBlock();
				}
				}
				setState(167);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(169);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__4) {
				{
				setState(168);
				defaultBlock();
				}
			}

			setState(171);
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
		enterRule(_localctx, 22, RULE_caseBlock);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(173);
			match(T__3);
			setState(174);
			expresion(0);
			setState(175);
			match(COLON);
			setState(179);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 2815334696288584L) != 0)) {
				{
				{
				setState(176);
				declaraciones();
				}
				}
				setState(181);
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
		enterRule(_localctx, 24, RULE_defaultBlock);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(182);
			match(T__4);
			setState(183);
			match(COLON);
			setState(187);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 2815334696288584L) != 0)) {
				{
				{
				setState(184);
				declaraciones();
				}
				}
				setState(189);
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
		public LlamadaFuncionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_llamadaFuncion; }
	}

	public final LlamadaFuncionContext llamadaFuncion() throws RecognitionException {
		LlamadaFuncionContext _localctx = new LlamadaFuncionContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_llamadaFuncion);
		int _la;
		try {
			setState(229);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INDEXOF:
				enterOuterAlt(_localctx, 1);
				{
				setState(190);
				match(INDEXOF);
				setState(191);
				match(LPAREN);
				setState(200);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 2815334427852800L) != 0)) {
					{
					setState(192);
					expresion(0);
					setState(197);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==COMMA) {
						{
						{
						setState(193);
						match(COMMA);
						setState(194);
						expresion(0);
						}
						}
						setState(199);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					}
				}

				setState(202);
				match(RPAREN);
				}
				break;
			case JOIN:
				enterOuterAlt(_localctx, 2);
				{
				setState(203);
				match(JOIN);
				setState(204);
				match(LPAREN);
				setState(213);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 2815334427852800L) != 0)) {
					{
					setState(205);
					expresion(0);
					setState(210);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==COMMA) {
						{
						{
						setState(206);
						match(COMMA);
						setState(207);
						expresion(0);
						}
						}
						setState(212);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					}
				}

				setState(215);
				match(RPAREN);
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 3);
				{
				setState(216);
				match(ID);
				setState(217);
				match(LPAREN);
				setState(226);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 2815334427852800L) != 0)) {
					{
					setState(218);
					expresion(0);
					setState(223);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==COMMA) {
						{
						{
						setState(219);
						match(COMMA);
						setState(220);
						expresion(0);
						}
						}
						setState(225);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					}
				}

				setState(228);
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
		enterRule(_localctx, 28, RULE_whileDcl);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(231);
			match(T__5);
			setState(232);
			match(LPAREN);
			setState(233);
			expresion(0);
			setState(234);
			match(RPAREN);
			setState(235);
			match(LBRACK);
			setState(239);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 2815334696288584L) != 0)) {
				{
				{
				setState(236);
				declaraciones();
				}
				}
				setState(241);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(242);
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
	public static class OrContext extends ExpresionContext {
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public TerminalNode OR() { return getToken(VlangParser.OR, 0); }
		public OrContext(ExpresionContext ctx) { copyFrom(ctx); }
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
	public static class SliceCreacionvContext extends ExpresionContext {
		public TipoSliceContext tipoSlice() {
			return getRuleContext(TipoSliceContext.class,0);
		}
		public TerminalNode LBRACE() { return getToken(VlangParser.LBRACE, 0); }
		public TerminalNode RBRACE() { return getToken(VlangParser.RBRACE, 0); }
		public ListaExpresionesContext listaExpresiones() {
			return getRuleContext(ListaExpresionesContext.class,0);
		}
		public SliceCreacionvContext(ExpresionContext ctx) { copyFrom(ctx); }
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
		int _startState = 30;
		enterRecursionRule(_localctx, 30, RULE_expresion, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(279);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,33,_ctx) ) {
			case 1:
				{
				_localctx = new ValorexprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(245);
				valor();
				}
				break;
			case 2:
				{
				_localctx = new ParentesisexpreContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(246);
				match(LPAREN);
				setState(247);
				expresion(0);
				setState(248);
				match(RPAREN);
				}
				break;
			case 3:
				{
				_localctx = new CorchetesexpreContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(250);
				match(LBRACK);
				setState(251);
				expresion(0);
				setState(252);
				match(RBRACK);
				}
				break;
			case 4:
				{
				_localctx = new SliceCreacionvContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(254);
				tipoSlice();
				setState(255);
				match(LBRACE);
				setState(257);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 2815334427852800L) != 0)) {
					{
					setState(256);
					listaExpresiones();
					}
				}

				setState(259);
				match(RBRACE);
				}
				break;
			case 5:
				{
				_localctx = new LlamadaFuncionExprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(261);
				llamadaFuncion();
				}
				break;
			case 6:
				{
				_localctx = new UnarioContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(262);
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
				setState(263);
				expresion(13);
				}
				break;
			case 7:
				{
				_localctx = new IdContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(264);
				match(ID);
				}
				break;
			case 8:
				{
				_localctx = new IncredecrContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(265);
				incredecre();
				}
				break;
			case 9:
				{
				_localctx = new Expdotexp1Context(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(266);
				match(ID);
				setState(267);
				match(DOT);
				setState(268);
				match(ID);
				}
				break;
			case 10:
				{
				_localctx = new ExpdotexpContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(269);
				match(ID);
				setState(270);
				match(DOT);
				setState(271);
				expresion(3);
				}
				break;
			case 11:
				{
				_localctx = new AsignacionLUEGOContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(272);
				match(ID);
				setState(273);
				match(TIPO);
				setState(274);
				match(ASSIGN);
				setState(275);
				expresion(2);
				}
				break;
			case 12:
				{
				_localctx = new IMCPLICITContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(276);
				match(ID);
				setState(277);
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
				setState(278);
				expresion(1);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(301);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,35,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(299);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,34,_ctx) ) {
					case 1:
						{
						_localctx = new OPERADORESLOGICOSContext(new ExpresionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(281);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(282);
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
						setState(283);
						expresion(13);
						}
						break;
					case 2:
						{
						_localctx = new MultdivmodContext(new ExpresionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(284);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(285);
						((MultdivmodContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 481036337152L) != 0)) ) {
							((MultdivmodContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(286);
						expresion(12);
						}
						break;
					case 3:
						{
						_localctx = new SumresContext(new ExpresionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(287);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(288);
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
						setState(289);
						expresion(11);
						}
						break;
					case 4:
						{
						_localctx = new RelacionalesContext(new ExpresionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(290);
						if (!(precpred(_ctx, 9))) throw new FailedPredicateException(this, "precpred(_ctx, 9)");
						setState(291);
						((RelacionalesContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 263882790666240L) != 0)) ) {
							((RelacionalesContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(292);
						expresion(10);
						}
						break;
					case 5:
						{
						_localctx = new IgualdadContext(new ExpresionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(293);
						if (!(precpred(_ctx, 8))) throw new FailedPredicateException(this, "precpred(_ctx, 8)");
						setState(294);
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
						setState(295);
						expresion(9);
						}
						break;
					case 6:
						{
						_localctx = new OrContext(new ExpresionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(296);
						if (!(precpred(_ctx, 7))) throw new FailedPredicateException(this, "precpred(_ctx, 7)");
						setState(297);
						match(OR);
						setState(298);
						expresion(8);
						}
						break;
					}
					} 
				}
				setState(303);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,35,_ctx);
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
		enterRule(_localctx, 32, RULE_parametros);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(304);
			expresion(0);
			setState(309);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(305);
				match(COMMA);
				setState(306);
				expresion(0);
				}
				}
				setState(311);
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
		enterRule(_localctx, 34, RULE_valores);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(312);
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
		enterRule(_localctx, 36, RULE_valor);
		try {
			setState(319);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case ENTERO:
				_localctx = new ValorEnteroContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(314);
				match(ENTERO);
				}
				break;
			case DECIMAL:
				_localctx = new ValorDecimalContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(315);
				match(DECIMAL);
				}
				break;
			case CADENA:
				_localctx = new ValorCadenaContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(316);
				match(CADENA);
				}
				break;
			case BOOLEANO:
				_localctx = new ValorBooleanoContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(317);
				match(BOOLEANO);
				}
				break;
			case CARACTER:
				_localctx = new ValorCaracterContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(318);
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
		enterRule(_localctx, 38, RULE_listaExpresiones);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(321);
			expresion(0);
			setState(326);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(322);
				match(COMMA);
				setState(323);
				expresion(0);
				}
				}
				setState(328);
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
		enterRule(_localctx, 40, RULE_incredecre);
		try {
			setState(333);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,39,_ctx) ) {
			case 1:
				_localctx = new IncrementoContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(329);
				match(ID);
				setState(330);
				match(INC);
				}
				break;
			case 2:
				_localctx = new DecrementoContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(331);
				match(ID);
				setState(332);
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
		case 15:
			return expresion_sempred((ExpresionContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expresion_sempred(ExpresionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 12);
		case 1:
			return precpred(_ctx, 11);
		case 2:
			return precpred(_ctx, 10);
		case 3:
			return precpred(_ctx, 9);
		case 4:
			return precpred(_ctx, 8);
		case 5:
			return precpred(_ctx, 7);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001=\u0150\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002\u0012\u0007\u0012"+
		"\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0001\u0000\u0005\u0000"+
		",\b\u0000\n\u0000\f\u0000/\t\u0000\u0001\u0000\u0001\u0000\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0002"+
		"\u0001\u0002\u0005\u0002;\b\u0002\n\u0002\f\u0002>\t\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0003\u0001\u0003\u0003\u0003D\b\u0003\u0001\u0004"+
		"\u0001\u0004\u0001\u0004\u0003\u0004I\b\u0004\u0001\u0004\u0001\u0004"+
		"\u0003\u0004M\b\u0004\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0005\u0006"+
		"X\b\u0006\n\u0006\f\u0006[\t\u0006\u0003\u0006]\b\u0006\u0001\u0006\u0001"+
		"\u0006\u0001\u0006\u0003\u0006b\b\u0006\u0001\u0007\u0001\u0007\u0001"+
		"\u0007\u0001\u0007\u0003\u0007h\b\u0007\u0001\b\u0001\b\u0003\bl\b\b\u0001"+
		"\b\u0001\b\u0003\bp\b\b\u0001\b\u0001\b\u0005\bt\b\b\n\b\f\bw\t\b\u0001"+
		"\b\u0001\b\u0001\b\u0001\b\u0005\b}\b\b\n\b\f\b\u0080\t\b\u0001\b\u0003"+
		"\b\u0083\b\b\u0001\t\u0001\t\u0003\t\u0087\b\t\u0001\t\u0003\t\u008a\b"+
		"\t\u0001\t\u0001\t\u0001\t\u0001\t\u0003\t\u0090\b\t\u0001\t\u0001\t\u0001"+
		"\t\u0001\t\u0003\t\u0096\b\t\u0001\t\u0001\t\u0003\t\u009a\b\t\u0001\t"+
		"\u0001\t\u0003\t\u009e\b\t\u0001\n\u0001\n\u0001\n\u0001\n\u0005\n\u00a4"+
		"\b\n\n\n\f\n\u00a7\t\n\u0001\n\u0003\n\u00aa\b\n\u0001\n\u0001\n\u0001"+
		"\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0005\u000b\u00b2\b\u000b\n"+
		"\u000b\f\u000b\u00b5\t\u000b\u0001\f\u0001\f\u0001\f\u0005\f\u00ba\b\f"+
		"\n\f\f\f\u00bd\t\f\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0005\r\u00c4"+
		"\b\r\n\r\f\r\u00c7\t\r\u0003\r\u00c9\b\r\u0001\r\u0001\r\u0001\r\u0001"+
		"\r\u0001\r\u0001\r\u0005\r\u00d1\b\r\n\r\f\r\u00d4\t\r\u0003\r\u00d6\b"+
		"\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0005\r\u00de\b\r\n"+
		"\r\f\r\u00e1\t\r\u0003\r\u00e3\b\r\u0001\r\u0003\r\u00e6\b\r\u0001\u000e"+
		"\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0005\u000e"+
		"\u00ee\b\u000e\n\u000e\f\u000e\u00f1\t\u000e\u0001\u000e\u0001\u000e\u0001"+
		"\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001"+
		"\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001"+
		"\u000f\u0003\u000f\u0102\b\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001"+
		"\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001"+
		"\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001"+
		"\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0003\u000f\u0118"+
		"\b\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001"+
		"\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001"+
		"\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0001"+
		"\u000f\u0005\u000f\u012c\b\u000f\n\u000f\f\u000f\u012f\t\u000f\u0001\u0010"+
		"\u0001\u0010\u0001\u0010\u0005\u0010\u0134\b\u0010\n\u0010\f\u0010\u0137"+
		"\t\u0010\u0001\u0011\u0001\u0011\u0001\u0012\u0001\u0012\u0001\u0012\u0001"+
		"\u0012\u0001\u0012\u0003\u0012\u0140\b\u0012\u0001\u0013\u0001\u0013\u0001"+
		"\u0013\u0005\u0013\u0145\b\u0013\n\u0013\f\u0013\u0148\t\u0013\u0001\u0014"+
		"\u0001\u0014\u0001\u0014\u0001\u0014\u0003\u0014\u014e\b\u0014\u0001\u0014"+
		"\u0000\u0001\u001e\u0015\u0000\u0002\u0004\u0006\b\n\f\u000e\u0010\u0012"+
		"\u0014\u0016\u0018\u001a\u001c\u001e \"$&(\u0000\u0007\u0002\u0000##\'"+
		"\'\u0001\u0000 !\u0001\u0000()\u0001\u0000$&\u0001\u0000\"#\u0001\u0000"+
		",/\u0001\u0000*+\u0177\u0000-\u0001\u0000\u0000\u0000\u00022\u0001\u0000"+
		"\u0000\u0000\u00048\u0001\u0000\u0000\u0000\u0006C\u0001\u0000\u0000\u0000"+
		"\bE\u0001\u0000\u0000\u0000\nN\u0001\u0000\u0000\u0000\fa\u0001\u0000"+
		"\u0000\u0000\u000eg\u0001\u0000\u0000\u0000\u0010i\u0001\u0000\u0000\u0000"+
		"\u0012\u009d\u0001\u0000\u0000\u0000\u0014\u009f\u0001\u0000\u0000\u0000"+
		"\u0016\u00ad\u0001\u0000\u0000\u0000\u0018\u00b6\u0001\u0000\u0000\u0000"+
		"\u001a\u00e5\u0001\u0000\u0000\u0000\u001c\u00e7\u0001\u0000\u0000\u0000"+
		"\u001e\u0117\u0001\u0000\u0000\u0000 \u0130\u0001\u0000\u0000\u0000\""+
		"\u0138\u0001\u0000\u0000\u0000$\u013f\u0001\u0000\u0000\u0000&\u0141\u0001"+
		"\u0000\u0000\u0000(\u014d\u0001\u0000\u0000\u0000*,\u0003\u0002\u0001"+
		"\u0000+*\u0001\u0000\u0000\u0000,/\u0001\u0000\u0000\u0000-+\u0001\u0000"+
		"\u0000\u0000-.\u0001\u0000\u0000\u0000.0\u0001\u0000\u0000\u0000/-\u0001"+
		"\u0000\u0000\u000001\u0005\u0000\u0000\u00011\u0001\u0001\u0000\u0000"+
		"\u000023\u0005\u0001\u0000\u000034\u0005\u0002\u0000\u000045\u00051\u0000"+
		"\u000056\u00052\u0000\u000067\u0003\u0004\u0002\u00007\u0003\u0001\u0000"+
		"\u0000\u00008<\u00055\u0000\u00009;\u0003\u0006\u0003\u0000:9\u0001\u0000"+
		"\u0000\u0000;>\u0001\u0000\u0000\u0000<:\u0001\u0000\u0000\u0000<=\u0001"+
		"\u0000\u0000\u0000=?\u0001\u0000\u0000\u0000><\u0001\u0000\u0000\u0000"+
		"?@\u00056\u0000\u0000@\u0005\u0001\u0000\u0000\u0000AD\u0003\b\u0004\u0000"+
		"BD\u0003\f\u0006\u0000CA\u0001\u0000\u0000\u0000CB\u0001\u0000\u0000\u0000"+
		"D\u0007\u0001\u0000\u0000\u0000EF\u0005\u0003\u0000\u0000FH\u0005\u001d"+
		"\u0000\u0000GI\u0005\u0007\u0000\u0000HG\u0001\u0000\u0000\u0000HI\u0001"+
		"\u0000\u0000\u0000IL\u0001\u0000\u0000\u0000JK\u00050\u0000\u0000KM\u0003"+
		"\u001e\u000f\u0000LJ\u0001\u0000\u0000\u0000LM\u0001\u0000\u0000\u0000"+
		"M\t\u0001\u0000\u0000\u0000NO\u00053\u0000\u0000OP\u00054\u0000\u0000"+
		"PQ\u0005\u0007\u0000\u0000Q\u000b\u0001\u0000\u0000\u0000RS\u0005\u001c"+
		"\u0000\u0000S\\\u00051\u0000\u0000TY\u0003\u001e\u000f\u0000UV\u0005:"+
		"\u0000\u0000VX\u0003\u001e\u000f\u0000WU\u0001\u0000\u0000\u0000X[\u0001"+
		"\u0000\u0000\u0000YW\u0001\u0000\u0000\u0000YZ\u0001\u0000\u0000\u0000"+
		"Z]\u0001\u0000\u0000\u0000[Y\u0001\u0000\u0000\u0000\\T\u0001\u0000\u0000"+
		"\u0000\\]\u0001\u0000\u0000\u0000]^\u0001\u0000\u0000\u0000^b\u00052\u0000"+
		"\u0000_b\u0003\u001e\u000f\u0000`b\u0003\u000e\u0007\u0000aR\u0001\u0000"+
		"\u0000\u0000a_\u0001\u0000\u0000\u0000a`\u0001\u0000\u0000\u0000b\r\u0001"+
		"\u0000\u0000\u0000ch\u0003\u0010\b\u0000dh\u0003\u0012\t\u0000eh\u0003"+
		"\u0014\n\u0000fh\u0003\u001c\u000e\u0000gc\u0001\u0000\u0000\u0000gd\u0001"+
		"\u0000\u0000\u0000ge\u0001\u0000\u0000\u0000gf\u0001\u0000\u0000\u0000"+
		"h\u000f\u0001\u0000\u0000\u0000ik\u0005\b\u0000\u0000jl\u00051\u0000\u0000"+
		"kj\u0001\u0000\u0000\u0000kl\u0001\u0000\u0000\u0000lm\u0001\u0000\u0000"+
		"\u0000mo\u0003\u001e\u000f\u0000np\u00052\u0000\u0000on\u0001\u0000\u0000"+
		"\u0000op\u0001\u0000\u0000\u0000pq\u0001\u0000\u0000\u0000qu\u00055\u0000"+
		"\u0000rt\u0003\u0006\u0003\u0000sr\u0001\u0000\u0000\u0000tw\u0001\u0000"+
		"\u0000\u0000us\u0001\u0000\u0000\u0000uv\u0001\u0000\u0000\u0000vx\u0001"+
		"\u0000\u0000\u0000wu\u0001\u0000\u0000\u0000x\u0082\u00056\u0000\u0000"+
		"yz\u0005\b\u0000\u0000z~\u00055\u0000\u0000{}\u0003\u0006\u0003\u0000"+
		"|{\u0001\u0000\u0000\u0000}\u0080\u0001\u0000\u0000\u0000~|\u0001\u0000"+
		"\u0000\u0000~\u007f\u0001\u0000\u0000\u0000\u007f\u0081\u0001\u0000\u0000"+
		"\u0000\u0080~\u0001\u0000\u0000\u0000\u0081\u0083\u00056\u0000\u0000\u0082"+
		"y\u0001\u0000\u0000\u0000\u0082\u0083\u0001\u0000\u0000\u0000\u0083\u0011"+
		"\u0001\u0000\u0000\u0000\u0084\u0086\u0005\b\u0000\u0000\u0085\u0087\u0005"+
		"1\u0000\u0000\u0086\u0085\u0001\u0000\u0000\u0000\u0086\u0087\u0001\u0000"+
		"\u0000\u0000\u0087\u0089\u0001\u0000\u0000\u0000\u0088\u008a\u0003\f\u0006"+
		"\u0000\u0089\u0088\u0001\u0000\u0000\u0000\u0089\u008a\u0001\u0000\u0000"+
		"\u0000\u008a\u008b\u0001\u0000\u0000\u0000\u008b\u008c\u00057\u0000\u0000"+
		"\u008c\u008d\u0003\u001e\u000f\u0000\u008d\u008f\u00057\u0000\u0000\u008e"+
		"\u0090\u0003\f\u0006\u0000\u008f\u008e\u0001\u0000\u0000\u0000\u008f\u0090"+
		"\u0001\u0000\u0000\u0000\u0090\u0091\u0001\u0000\u0000\u0000\u0091\u0092"+
		"\u0003\u0004\u0002\u0000\u0092\u009e\u0001\u0000\u0000\u0000\u0093\u0095"+
		"\u0005\b\u0000\u0000\u0094\u0096\u00051\u0000\u0000\u0095\u0094\u0001"+
		"\u0000\u0000\u0000\u0095\u0096\u0001\u0000\u0000\u0000\u0096\u0097\u0001"+
		"\u0000\u0000\u0000\u0097\u0099\u0003\u001e\u000f\u0000\u0098\u009a\u0005"+
		"2\u0000\u0000\u0099\u0098\u0001\u0000\u0000\u0000\u0099\u009a\u0001\u0000"+
		"\u0000\u0000\u009a\u009b\u0001\u0000\u0000\u0000\u009b\u009c\u0003\u0004"+
		"\u0002\u0000\u009c\u009e\u0001\u0000\u0000\u0000\u009d\u0084\u0001\u0000"+
		"\u0000\u0000\u009d\u0093\u0001\u0000\u0000\u0000\u009e\u0013\u0001\u0000"+
		"\u0000\u0000\u009f\u00a0\u0005\b\u0000\u0000\u00a0\u00a1\u0003\u001e\u000f"+
		"\u0000\u00a1\u00a5\u00055\u0000\u0000\u00a2\u00a4\u0003\u0016\u000b\u0000"+
		"\u00a3\u00a2\u0001\u0000\u0000\u0000\u00a4\u00a7\u0001\u0000\u0000\u0000"+
		"\u00a5\u00a3\u0001\u0000\u0000\u0000\u00a5\u00a6\u0001\u0000\u0000\u0000"+
		"\u00a6\u00a9\u0001\u0000\u0000\u0000\u00a7\u00a5\u0001\u0000\u0000\u0000"+
		"\u00a8\u00aa\u0003\u0018\f\u0000\u00a9\u00a8\u0001\u0000\u0000\u0000\u00a9"+
		"\u00aa\u0001\u0000\u0000\u0000\u00aa\u00ab\u0001\u0000\u0000\u0000\u00ab"+
		"\u00ac\u00056\u0000\u0000\u00ac\u0015\u0001\u0000\u0000\u0000\u00ad\u00ae"+
		"\u0005\u0004\u0000\u0000\u00ae\u00af\u0003\u001e\u000f\u0000\u00af\u00b3"+
		"\u00058\u0000\u0000\u00b0\u00b2\u0003\u0006\u0003\u0000\u00b1\u00b0\u0001"+
		"\u0000\u0000\u0000\u00b2\u00b5\u0001\u0000\u0000\u0000\u00b3\u00b1\u0001"+
		"\u0000\u0000\u0000\u00b3\u00b4\u0001\u0000\u0000\u0000\u00b4\u0017\u0001"+
		"\u0000\u0000\u0000\u00b5\u00b3\u0001\u0000\u0000\u0000\u00b6\u00b7\u0005"+
		"\u0005\u0000\u0000\u00b7\u00bb\u00058\u0000\u0000\u00b8\u00ba\u0003\u0006"+
		"\u0003\u0000\u00b9\u00b8\u0001\u0000\u0000\u0000\u00ba\u00bd\u0001\u0000"+
		"\u0000\u0000\u00bb\u00b9\u0001\u0000\u0000\u0000\u00bb\u00bc\u0001\u0000"+
		"\u0000\u0000\u00bc\u0019\u0001\u0000\u0000\u0000\u00bd\u00bb\u0001\u0000"+
		"\u0000\u0000\u00be\u00bf\u0005\u0010\u0000\u0000\u00bf\u00c8\u00051\u0000"+
		"\u0000\u00c0\u00c5\u0003\u001e\u000f\u0000\u00c1\u00c2\u0005:\u0000\u0000"+
		"\u00c2\u00c4\u0003\u001e\u000f\u0000\u00c3\u00c1\u0001\u0000\u0000\u0000"+
		"\u00c4\u00c7\u0001\u0000\u0000\u0000\u00c5\u00c3\u0001\u0000\u0000\u0000"+
		"\u00c5\u00c6\u0001\u0000\u0000\u0000\u00c6\u00c9\u0001\u0000\u0000\u0000"+
		"\u00c7\u00c5\u0001\u0000\u0000\u0000\u00c8\u00c0\u0001\u0000\u0000\u0000"+
		"\u00c8\u00c9\u0001\u0000\u0000\u0000\u00c9\u00ca\u0001\u0000\u0000\u0000"+
		"\u00ca\u00e6\u00052\u0000\u0000\u00cb\u00cc\u0005\u0011\u0000\u0000\u00cc"+
		"\u00d5\u00051\u0000\u0000\u00cd\u00d2\u0003\u001e\u000f\u0000\u00ce\u00cf"+
		"\u0005:\u0000\u0000\u00cf\u00d1\u0003\u001e\u000f\u0000\u00d0\u00ce\u0001"+
		"\u0000\u0000\u0000\u00d1\u00d4\u0001\u0000\u0000\u0000\u00d2\u00d0\u0001"+
		"\u0000\u0000\u0000\u00d2\u00d3\u0001\u0000\u0000\u0000\u00d3\u00d6\u0001"+
		"\u0000\u0000\u0000\u00d4\u00d2\u0001\u0000\u0000\u0000\u00d5\u00cd\u0001"+
		"\u0000\u0000\u0000\u00d5\u00d6\u0001\u0000\u0000\u0000\u00d6\u00d7\u0001"+
		"\u0000\u0000\u0000\u00d7\u00e6\u00052\u0000\u0000\u00d8\u00d9\u0005\u001d"+
		"\u0000\u0000\u00d9\u00e2\u00051\u0000\u0000\u00da\u00df\u0003\u001e\u000f"+
		"\u0000\u00db\u00dc\u0005:\u0000\u0000\u00dc\u00de\u0003\u001e\u000f\u0000"+
		"\u00dd\u00db\u0001\u0000\u0000\u0000\u00de\u00e1\u0001\u0000\u0000\u0000"+
		"\u00df\u00dd\u0001\u0000\u0000\u0000\u00df\u00e0\u0001\u0000\u0000\u0000"+
		"\u00e0\u00e3\u0001\u0000\u0000\u0000\u00e1\u00df\u0001\u0000\u0000\u0000"+
		"\u00e2\u00da\u0001\u0000\u0000\u0000\u00e2\u00e3\u0001\u0000\u0000\u0000"+
		"\u00e3\u00e4\u0001\u0000\u0000\u0000\u00e4\u00e6\u00052\u0000\u0000\u00e5"+
		"\u00be\u0001\u0000\u0000\u0000\u00e5\u00cb\u0001\u0000\u0000\u0000\u00e5"+
		"\u00d8\u0001\u0000\u0000\u0000\u00e6\u001b\u0001\u0000\u0000\u0000\u00e7"+
		"\u00e8\u0005\u0006\u0000\u0000\u00e8\u00e9\u00051\u0000\u0000\u00e9\u00ea"+
		"\u0003\u001e\u000f\u0000\u00ea\u00eb\u00052\u0000\u0000\u00eb\u00ef\u0005"+
		"3\u0000\u0000\u00ec\u00ee\u0003\u0006\u0003\u0000\u00ed\u00ec\u0001\u0000"+
		"\u0000\u0000\u00ee\u00f1\u0001\u0000\u0000\u0000\u00ef\u00ed\u0001\u0000"+
		"\u0000\u0000\u00ef\u00f0\u0001\u0000\u0000\u0000\u00f0\u00f2\u0001\u0000"+
		"\u0000\u0000\u00f1\u00ef\u0001\u0000\u0000\u0000\u00f2\u00f3\u00054\u0000"+
		"\u0000\u00f3\u001d\u0001\u0000\u0000\u0000\u00f4\u00f5\u0006\u000f\uffff"+
		"\uffff\u0000\u00f5\u0118\u0003$\u0012\u0000\u00f6\u00f7\u00051\u0000\u0000"+
		"\u00f7\u00f8\u0003\u001e\u000f\u0000\u00f8\u00f9\u00052\u0000\u0000\u00f9"+
		"\u0118\u0001\u0000\u0000\u0000\u00fa\u00fb\u00053\u0000\u0000\u00fb\u00fc"+
		"\u0003\u001e\u000f\u0000\u00fc\u00fd\u00054\u0000\u0000\u00fd\u0118\u0001"+
		"\u0000\u0000\u0000\u00fe\u00ff\u0003\n\u0005\u0000\u00ff\u0101\u00055"+
		"\u0000\u0000\u0100\u0102\u0003&\u0013\u0000\u0101\u0100\u0001\u0000\u0000"+
		"\u0000\u0101\u0102\u0001\u0000\u0000\u0000\u0102\u0103\u0001\u0000\u0000"+
		"\u0000\u0103\u0104\u00056\u0000\u0000\u0104\u0118\u0001\u0000\u0000\u0000"+
		"\u0105\u0118\u0003\u001a\r\u0000\u0106\u0107\u0007\u0000\u0000\u0000\u0107"+
		"\u0118\u0003\u001e\u000f\r\u0108\u0118\u0005\u001d\u0000\u0000\u0109\u0118"+
		"\u0003(\u0014\u0000\u010a\u010b\u0005\u001d\u0000\u0000\u010b\u010c\u0005"+
		"9\u0000\u0000\u010c\u0118\u0005\u001d\u0000\u0000\u010d\u010e\u0005\u001d"+
		"\u0000\u0000\u010e\u010f\u00059\u0000\u0000\u010f\u0118\u0003\u001e\u000f"+
		"\u0003\u0110\u0111\u0005\u001d\u0000\u0000\u0111\u0112\u0005\u0007\u0000"+
		"\u0000\u0112\u0113\u00050\u0000\u0000\u0113\u0118\u0003\u001e\u000f\u0002"+
		"\u0114\u0115\u0005\u001d\u0000\u0000\u0115\u0116\u0007\u0001\u0000\u0000"+
		"\u0116\u0118\u0003\u001e\u000f\u0001\u0117\u00f4\u0001\u0000\u0000\u0000"+
		"\u0117\u00f6\u0001\u0000\u0000\u0000\u0117\u00fa\u0001\u0000\u0000\u0000"+
		"\u0117\u00fe\u0001\u0000\u0000\u0000\u0117\u0105\u0001\u0000\u0000\u0000"+
		"\u0117\u0106\u0001\u0000\u0000\u0000\u0117\u0108\u0001\u0000\u0000\u0000"+
		"\u0117\u0109\u0001\u0000\u0000\u0000\u0117\u010a\u0001\u0000\u0000\u0000"+
		"\u0117\u010d\u0001\u0000\u0000\u0000\u0117\u0110\u0001\u0000\u0000\u0000"+
		"\u0117\u0114\u0001\u0000\u0000\u0000\u0118\u012d\u0001\u0000\u0000\u0000"+
		"\u0119\u011a\n\f\u0000\u0000\u011a\u011b\u0007\u0002\u0000\u0000\u011b"+
		"\u012c\u0003\u001e\u000f\r\u011c\u011d\n\u000b\u0000\u0000\u011d\u011e"+
		"\u0007\u0003\u0000\u0000\u011e\u012c\u0003\u001e\u000f\f\u011f\u0120\n"+
		"\n\u0000\u0000\u0120\u0121\u0007\u0004\u0000\u0000\u0121\u012c\u0003\u001e"+
		"\u000f\u000b\u0122\u0123\n\t\u0000\u0000\u0123\u0124\u0007\u0005\u0000"+
		"\u0000\u0124\u012c\u0003\u001e\u000f\n\u0125\u0126\n\b\u0000\u0000\u0126"+
		"\u0127\u0007\u0006\u0000\u0000\u0127\u012c\u0003\u001e\u000f\t\u0128\u0129"+
		"\n\u0007\u0000\u0000\u0129\u012a\u0005(\u0000\u0000\u012a\u012c\u0003"+
		"\u001e\u000f\b\u012b\u0119\u0001\u0000\u0000\u0000\u012b\u011c\u0001\u0000"+
		"\u0000\u0000\u012b\u011f\u0001\u0000\u0000\u0000\u012b\u0122\u0001\u0000"+
		"\u0000\u0000\u012b\u0125\u0001\u0000\u0000\u0000\u012b\u0128\u0001\u0000"+
		"\u0000\u0000\u012c\u012f\u0001\u0000\u0000\u0000\u012d\u012b\u0001\u0000"+
		"\u0000\u0000\u012d\u012e\u0001\u0000\u0000\u0000\u012e\u001f\u0001\u0000"+
		"\u0000\u0000\u012f\u012d\u0001\u0000\u0000\u0000\u0130\u0135\u0003\u001e"+
		"\u000f\u0000\u0131\u0132\u0005:\u0000\u0000\u0132\u0134\u0003\u001e\u000f"+
		"\u0000\u0133\u0131\u0001\u0000\u0000\u0000\u0134\u0137\u0001\u0000\u0000"+
		"\u0000\u0135\u0133\u0001\u0000\u0000\u0000\u0135\u0136\u0001\u0000\u0000"+
		"\u0000\u0136!\u0001\u0000\u0000\u0000\u0137\u0135\u0001\u0000\u0000\u0000"+
		"\u0138\u0139\u0003$\u0012\u0000\u0139#\u0001\u0000\u0000\u0000\u013a\u0140"+
		"\u0005\u0013\u0000\u0000\u013b\u0140\u0005\u0014\u0000\u0000\u013c\u0140"+
		"\u0005\u0015\u0000\u0000\u013d\u0140\u0005\u0012\u0000\u0000\u013e\u0140"+
		"\u0005\u0016\u0000\u0000\u013f\u013a\u0001\u0000\u0000\u0000\u013f\u013b"+
		"\u0001\u0000\u0000\u0000\u013f\u013c\u0001\u0000\u0000\u0000\u013f\u013d"+
		"\u0001\u0000\u0000\u0000\u013f\u013e\u0001\u0000\u0000\u0000\u0140%\u0001"+
		"\u0000\u0000\u0000\u0141\u0146\u0003\u001e\u000f\u0000\u0142\u0143\u0005"+
		":\u0000\u0000\u0143\u0145\u0003\u001e\u000f\u0000\u0144\u0142\u0001\u0000"+
		"\u0000\u0000\u0145\u0148\u0001\u0000\u0000\u0000\u0146\u0144\u0001\u0000"+
		"\u0000\u0000\u0146\u0147\u0001\u0000\u0000\u0000\u0147\'\u0001\u0000\u0000"+
		"\u0000\u0148\u0146\u0001\u0000\u0000\u0000\u0149\u014a\u0005\u001d\u0000"+
		"\u0000\u014a\u014e\u0005\u001e\u0000\u0000\u014b\u014c\u0005\u001d\u0000"+
		"\u0000\u014c\u014e\u0005\u001f\u0000\u0000\u014d\u0149\u0001\u0000\u0000"+
		"\u0000\u014d\u014b\u0001\u0000\u0000\u0000\u014e)\u0001\u0000\u0000\u0000"+
		"(-<CHLY\\agkou~\u0082\u0086\u0089\u008f\u0095\u0099\u009d\u00a5\u00a9"+
		"\u00b3\u00bb\u00c5\u00c8\u00d2\u00d5\u00df\u00e2\u00e5\u00ef\u0101\u0117"+
		"\u012b\u012d\u0135\u013f\u0146\u014d";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}