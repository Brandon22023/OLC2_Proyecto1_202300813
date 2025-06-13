// Generated from /home/brandon/Escritorio/OLC2_Proyecto1_202300813/Backend/compiler/Vlang.g4 by ANTLR 4.13.1
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
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, FUNCIONES_PARECEN_CASTEO=7, 
		LEN=8, CAP=9, APPEND=10, BOOLEANO=11, ENTERO=12, DECIMAL=13, CADENA=14, 
		CARACTER=15, ATOI=16, PARSEFLOAT=17, TYPEOF=18, ID=19, PLUS=20, MINUS=21, 
		MUL=22, DIV=23, MOD=24, NOT=25, OR=26, EQ=27, NEQ=28, LT=29, LE=30, GT=31, 
		GE=32, ASSIGN=33, INC=34, DEC=35, LPAREN=36, RPAREN=37, LBRACK=38, RBRACK=39, 
		DOT=40, COMMA=41, WS=42, LINE_COMMENT=43, BLOCK_COMMENT=44;
	public static final int
		RULE_programa = 0, RULE_declaraciones = 1, RULE_varDcl = 2, RULE_stmt = 3, 
		RULE_sentencias_control = 4, RULE_ifDcl = 5, RULE_forDcl = 6, RULE_whileDcl = 7, 
		RULE_expresion = 8, RULE_parametros = 9, RULE_valores = 10, RULE_valor = 11, 
		RULE_incredecre = 12;
	private static String[] makeRuleNames() {
		return new String[] {
			"programa", "declaraciones", "varDcl", "stmt", "sentencias_control", 
			"ifDcl", "forDcl", "whileDcl", "expresion", "parametros", "valores", 
			"valor", "incredecre"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'mut'", "'print('", "'if'", "'else'", "'for'", "'while'", null, 
			"'len'", "'cap'", "'append'", null, null, null, null, null, "'Atoi'", 
			"'parseFloat'", "'typeOf'", null, "'+'", "'-'", "'*'", "'/'", "'%'", 
			"'!'", "'||'", "'=='", "'!='", "'<'", "'<='", "'>'", "'>='", "':='", 
			"'++'", "'--'", "'('", "')'", "'['", "']'", "'.'", "','"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, "FUNCIONES_PARECEN_CASTEO", 
			"LEN", "CAP", "APPEND", "BOOLEANO", "ENTERO", "DECIMAL", "CADENA", "CARACTER", 
			"ATOI", "PARSEFLOAT", "TYPEOF", "ID", "PLUS", "MINUS", "MUL", "DIV", 
			"MOD", "NOT", "OR", "EQ", "NEQ", "LT", "LE", "GT", "GE", "ASSIGN", "INC", 
			"DEC", "LPAREN", "RPAREN", "LBRACK", "RBRACK", "DOT", "COMMA", "WS", 
			"LINE_COMMENT", "BLOCK_COMMENT"
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
		public List<DeclaracionesContext> declaraciones() {
			return getRuleContexts(DeclaracionesContext.class);
		}
		public DeclaracionesContext declaraciones(int i) {
			return getRuleContext(DeclaracionesContext.class,i);
		}
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode EOF() { return getToken(VlangParser.EOF, 0); }
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
			setState(35);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(29);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 343633885294L) != 0)) {
					{
					{
					setState(26);
					declaraciones();
					}
					}
					setState(31);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(32);
				expresion(0);
				setState(33);
				match(EOF);
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
		enterRule(_localctx, 2, RULE_declaraciones);
		try {
			setState(39);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__0:
				enterOuterAlt(_localctx, 1);
				{
				setState(37);
				varDcl();
				}
				break;
			case T__1:
			case T__2:
			case T__4:
			case T__5:
			case BOOLEANO:
			case ENTERO:
			case DECIMAL:
			case CADENA:
			case CARACTER:
			case TYPEOF:
			case ID:
			case MINUS:
			case NOT:
			case LPAREN:
			case LBRACK:
				enterOuterAlt(_localctx, 2);
				{
				setState(38);
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
		public TerminalNode ASSIGN() { return getToken(VlangParser.ASSIGN, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public VariableDeclarationContext(VarDclContext ctx) { copyFrom(ctx); }
	}

	public final VarDclContext varDcl() throws RecognitionException {
		VarDclContext _localctx = new VarDclContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_varDcl);
		int _la;
		try {
			_localctx = new VariableDeclarationContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(41);
			match(T__0);
			setState(42);
			match(ID);
			setState(45);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ASSIGN) {
				{
				setState(43);
				match(ASSIGN);
				setState(44);
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
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(VlangParser.RPAREN, 0); }
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
		enterRule(_localctx, 6, RULE_stmt);
		try {
			setState(53);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case BOOLEANO:
			case ENTERO:
			case DECIMAL:
			case CADENA:
			case CARACTER:
			case TYPEOF:
			case ID:
			case MINUS:
			case NOT:
			case LPAREN:
			case LBRACK:
				_localctx = new ExpresionStatementContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(47);
				expresion(0);
				}
				break;
			case T__1:
				_localctx = new PrintStatementContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(48);
				match(T__1);
				setState(49);
				expresion(0);
				setState(50);
				match(RPAREN);
				}
				break;
			case T__2:
			case T__4:
			case T__5:
				_localctx = new ControlStatementContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(52);
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
		enterRule(_localctx, 8, RULE_sentencias_control);
		try {
			setState(58);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case T__2:
				_localctx = new If_contextContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(55);
				ifDcl();
				}
				break;
			case T__4:
				_localctx = new For_contextContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(56);
				forDcl();
				}
				break;
			case T__5:
				_localctx = new While_contextContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(57);
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
	public static class IfDclContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(VlangParser.LPAREN, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(VlangParser.RPAREN, 0); }
		public List<TerminalNode> LBRACK() { return getTokens(VlangParser.LBRACK); }
		public TerminalNode LBRACK(int i) {
			return getToken(VlangParser.LBRACK, i);
		}
		public List<TerminalNode> RBRACK() { return getTokens(VlangParser.RBRACK); }
		public TerminalNode RBRACK(int i) {
			return getToken(VlangParser.RBRACK, i);
		}
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
		enterRule(_localctx, 10, RULE_ifDcl);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(60);
			match(T__2);
			setState(61);
			match(LPAREN);
			setState(62);
			expresion(0);
			setState(63);
			match(RPAREN);
			setState(64);
			match(LBRACK);
			setState(68);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 343633885294L) != 0)) {
				{
				{
				setState(65);
				declaraciones();
				}
				}
				setState(70);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(71);
			match(RBRACK);
			setState(81);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__3) {
				{
				setState(72);
				match(T__3);
				setState(73);
				match(LBRACK);
				setState(77);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 343633885294L) != 0)) {
					{
					{
					setState(74);
					declaraciones();
					}
					}
					setState(79);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(80);
				match(RBRACK);
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
	public static class ForDclContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(VlangParser.LPAREN, 0); }
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public TerminalNode ASSIGN() { return getToken(VlangParser.ASSIGN, 0); }
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public TerminalNode COMMA() { return getToken(VlangParser.COMMA, 0); }
		public TerminalNode RPAREN() { return getToken(VlangParser.RPAREN, 0); }
		public TerminalNode LBRACK() { return getToken(VlangParser.LBRACK, 0); }
		public TerminalNode RBRACK() { return getToken(VlangParser.RBRACK, 0); }
		public List<DeclaracionesContext> declaraciones() {
			return getRuleContexts(DeclaracionesContext.class);
		}
		public DeclaracionesContext declaraciones(int i) {
			return getRuleContext(DeclaracionesContext.class,i);
		}
		public ForDclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_forDcl; }
	}

	public final ForDclContext forDcl() throws RecognitionException {
		ForDclContext _localctx = new ForDclContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_forDcl);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(83);
			match(T__4);
			setState(84);
			match(LPAREN);
			setState(85);
			match(ID);
			setState(86);
			match(ASSIGN);
			setState(87);
			expresion(0);
			setState(88);
			match(COMMA);
			setState(89);
			expresion(0);
			setState(90);
			match(RPAREN);
			setState(91);
			match(LBRACK);
			setState(95);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 343633885294L) != 0)) {
				{
				{
				setState(92);
				declaraciones();
				}
				}
				setState(97);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(98);
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
		enterRule(_localctx, 14, RULE_whileDcl);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(100);
			match(T__5);
			setState(101);
			match(LPAREN);
			setState(102);
			expresion(0);
			setState(103);
			match(RPAREN);
			setState(104);
			match(LBRACK);
			setState(108);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 343633885294L) != 0)) {
				{
				{
				setState(105);
				declaraciones();
				}
				}
				setState(110);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(111);
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
	public static class Casteo_paratipoContext extends ExpresionContext {
		public TerminalNode TYPEOF() { return getToken(VlangParser.TYPEOF, 0); }
		public TerminalNode LPAREN() { return getToken(VlangParser.LPAREN, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(VlangParser.RPAREN, 0); }
		public Casteo_paratipoContext(ExpresionContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IncredecrContext extends ExpresionContext {
		public IncredecreContext incredecre() {
			return getRuleContext(IncredecreContext.class,0);
		}
		public IncredecrContext(ExpresionContext ctx) { copyFrom(ctx); }
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
	public static class AsignacionforContext extends ExpresionContext {
		public TerminalNode ID() { return getToken(VlangParser.ID, 0); }
		public TerminalNode ASSIGN() { return getToken(VlangParser.ASSIGN, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public AsignacionforContext(ExpresionContext ctx) { copyFrom(ctx); }
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
		int _startState = 16;
		enterRecursionRule(_localctx, 16, RULE_expresion, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(141);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,11,_ctx) ) {
			case 1:
				{
				_localctx = new ValorexprContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(114);
				valor();
				}
				break;
			case 2:
				{
				_localctx = new ParentesisexpreContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(115);
				match(LPAREN);
				setState(116);
				expresion(0);
				setState(117);
				match(RPAREN);
				}
				break;
			case 3:
				{
				_localctx = new CorchetesexpreContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(119);
				match(LBRACK);
				setState(120);
				expresion(0);
				setState(121);
				match(RBRACK);
				}
				break;
			case 4:
				{
				_localctx = new UnarioContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(123);
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
				setState(124);
				expresion(12);
				}
				break;
			case 5:
				{
				_localctx = new IdContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(125);
				match(ID);
				}
				break;
			case 6:
				{
				_localctx = new IncredecrContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(126);
				incredecre();
				}
				break;
			case 7:
				{
				_localctx = new Expdotexp1Context(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(127);
				match(ID);
				setState(128);
				match(DOT);
				setState(129);
				match(ID);
				}
				break;
			case 8:
				{
				_localctx = new ExpdotexpContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(130);
				match(ID);
				setState(131);
				match(DOT);
				setState(132);
				expresion(3);
				}
				break;
			case 9:
				{
				_localctx = new AsignacionforContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(133);
				match(ID);
				setState(134);
				match(ASSIGN);
				setState(135);
				expresion(2);
				}
				break;
			case 10:
				{
				_localctx = new Casteo_paratipoContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(136);
				match(TYPEOF);
				setState(137);
				match(LPAREN);
				setState(138);
				expresion(0);
				setState(139);
				match(RPAREN);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(160);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,13,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(158);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
					case 1:
						{
						_localctx = new MultdivmodContext(new ExpresionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(143);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(144);
						((MultdivmodContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 29360128L) != 0)) ) {
							((MultdivmodContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(145);
						expresion(12);
						}
						break;
					case 2:
						{
						_localctx = new SumresContext(new ExpresionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(146);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(147);
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
						setState(148);
						expresion(11);
						}
						break;
					case 3:
						{
						_localctx = new RelacionalesContext(new ExpresionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(149);
						if (!(precpred(_ctx, 9))) throw new FailedPredicateException(this, "precpred(_ctx, 9)");
						setState(150);
						((RelacionalesContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 8053063680L) != 0)) ) {
							((RelacionalesContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(151);
						expresion(10);
						}
						break;
					case 4:
						{
						_localctx = new IgualdadContext(new ExpresionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(152);
						if (!(precpred(_ctx, 8))) throw new FailedPredicateException(this, "precpred(_ctx, 8)");
						setState(153);
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
						setState(154);
						expresion(9);
						}
						break;
					case 5:
						{
						_localctx = new OrContext(new ExpresionContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(155);
						if (!(precpred(_ctx, 7))) throw new FailedPredicateException(this, "precpred(_ctx, 7)");
						setState(156);
						match(OR);
						setState(157);
						expresion(8);
						}
						break;
					}
					} 
				}
				setState(162);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,13,_ctx);
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
		enterRule(_localctx, 18, RULE_parametros);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(163);
			expresion(0);
			setState(168);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(164);
				match(COMMA);
				setState(165);
				expresion(0);
				}
				}
				setState(170);
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
		enterRule(_localctx, 20, RULE_valores);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(171);
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
		enterRule(_localctx, 22, RULE_valor);
		try {
			setState(178);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case ENTERO:
				_localctx = new ValorEnteroContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(173);
				match(ENTERO);
				}
				break;
			case DECIMAL:
				_localctx = new ValorDecimalContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(174);
				match(DECIMAL);
				}
				break;
			case CADENA:
				_localctx = new ValorCadenaContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(175);
				match(CADENA);
				}
				break;
			case BOOLEANO:
				_localctx = new ValorBooleanoContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(176);
				match(BOOLEANO);
				}
				break;
			case CARACTER:
				_localctx = new ValorCaracterContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(177);
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
		enterRule(_localctx, 24, RULE_incredecre);
		try {
			setState(184);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,16,_ctx) ) {
			case 1:
				_localctx = new IncrementoContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(180);
				match(ID);
				setState(181);
				match(INC);
				}
				break;
			case 2:
				_localctx = new DecrementoContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(182);
				match(ID);
				setState(183);
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
		case 8:
			return expresion_sempred((ExpresionContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expresion_sempred(ExpresionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 11);
		case 1:
			return precpred(_ctx, 10);
		case 2:
			return precpred(_ctx, 9);
		case 3:
			return precpred(_ctx, 8);
		case 4:
			return precpred(_ctx, 7);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001,\u00bb\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0001\u0000\u0005\u0000\u001c\b\u0000\n\u0000\f\u0000\u001f"+
		"\t\u0000\u0001\u0000\u0001\u0000\u0001\u0000\u0003\u0000$\b\u0000\u0001"+
		"\u0001\u0001\u0001\u0003\u0001(\b\u0001\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0003\u0002.\b\u0002\u0001\u0003\u0001\u0003\u0001"+
		"\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0003\u00036\b\u0003\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0003\u0004;\b\u0004\u0001\u0005\u0001"+
		"\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0005\u0005C\b"+
		"\u0005\n\u0005\f\u0005F\t\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0005\u0005\u0005L\b\u0005\n\u0005\f\u0005O\t\u0005\u0001\u0005\u0003"+
		"\u0005R\b\u0005\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001"+
		"\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0005"+
		"\u0006^\b\u0006\n\u0006\f\u0006a\t\u0006\u0001\u0006\u0001\u0006\u0001"+
		"\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0005"+
		"\u0007k\b\u0007\n\u0007\f\u0007n\t\u0007\u0001\u0007\u0001\u0007\u0001"+
		"\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001"+
		"\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001"+
		"\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001"+
		"\b\u0003\b\u008e\b\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001"+
		"\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0005"+
		"\b\u009f\b\b\n\b\f\b\u00a2\t\b\u0001\t\u0001\t\u0001\t\u0005\t\u00a7\b"+
		"\t\n\t\f\t\u00aa\t\t\u0001\n\u0001\n\u0001\u000b\u0001\u000b\u0001\u000b"+
		"\u0001\u000b\u0001\u000b\u0003\u000b\u00b3\b\u000b\u0001\f\u0001\f\u0001"+
		"\f\u0001\f\u0003\f\u00b9\b\f\u0001\f\u0000\u0001\u0010\r\u0000\u0002\u0004"+
		"\u0006\b\n\f\u000e\u0010\u0012\u0014\u0016\u0018\u0000\u0005\u0002\u0000"+
		"\u0015\u0015\u0019\u0019\u0001\u0000\u0016\u0018\u0001\u0000\u0014\u0015"+
		"\u0001\u0000\u001d \u0001\u0000\u001b\u001c\u00ce\u0000#\u0001\u0000\u0000"+
		"\u0000\u0002\'\u0001\u0000\u0000\u0000\u0004)\u0001\u0000\u0000\u0000"+
		"\u00065\u0001\u0000\u0000\u0000\b:\u0001\u0000\u0000\u0000\n<\u0001\u0000"+
		"\u0000\u0000\fS\u0001\u0000\u0000\u0000\u000ed\u0001\u0000\u0000\u0000"+
		"\u0010\u008d\u0001\u0000\u0000\u0000\u0012\u00a3\u0001\u0000\u0000\u0000"+
		"\u0014\u00ab\u0001\u0000\u0000\u0000\u0016\u00b2\u0001\u0000\u0000\u0000"+
		"\u0018\u00b8\u0001\u0000\u0000\u0000\u001a\u001c\u0003\u0002\u0001\u0000"+
		"\u001b\u001a\u0001\u0000\u0000\u0000\u001c\u001f\u0001\u0000\u0000\u0000"+
		"\u001d\u001b\u0001\u0000\u0000\u0000\u001d\u001e\u0001\u0000\u0000\u0000"+
		"\u001e$\u0001\u0000\u0000\u0000\u001f\u001d\u0001\u0000\u0000\u0000 !"+
		"\u0003\u0010\b\u0000!\"\u0005\u0000\u0000\u0001\"$\u0001\u0000\u0000\u0000"+
		"#\u001d\u0001\u0000\u0000\u0000# \u0001\u0000\u0000\u0000$\u0001\u0001"+
		"\u0000\u0000\u0000%(\u0003\u0004\u0002\u0000&(\u0003\u0006\u0003\u0000"+
		"\'%\u0001\u0000\u0000\u0000\'&\u0001\u0000\u0000\u0000(\u0003\u0001\u0000"+
		"\u0000\u0000)*\u0005\u0001\u0000\u0000*-\u0005\u0013\u0000\u0000+,\u0005"+
		"!\u0000\u0000,.\u0003\u0010\b\u0000-+\u0001\u0000\u0000\u0000-.\u0001"+
		"\u0000\u0000\u0000.\u0005\u0001\u0000\u0000\u0000/6\u0003\u0010\b\u0000"+
		"01\u0005\u0002\u0000\u000012\u0003\u0010\b\u000023\u0005%\u0000\u0000"+
		"36\u0001\u0000\u0000\u000046\u0003\b\u0004\u00005/\u0001\u0000\u0000\u0000"+
		"50\u0001\u0000\u0000\u000054\u0001\u0000\u0000\u00006\u0007\u0001\u0000"+
		"\u0000\u00007;\u0003\n\u0005\u00008;\u0003\f\u0006\u00009;\u0003\u000e"+
		"\u0007\u0000:7\u0001\u0000\u0000\u0000:8\u0001\u0000\u0000\u0000:9\u0001"+
		"\u0000\u0000\u0000;\t\u0001\u0000\u0000\u0000<=\u0005\u0003\u0000\u0000"+
		"=>\u0005$\u0000\u0000>?\u0003\u0010\b\u0000?@\u0005%\u0000\u0000@D\u0005"+
		"&\u0000\u0000AC\u0003\u0002\u0001\u0000BA\u0001\u0000\u0000\u0000CF\u0001"+
		"\u0000\u0000\u0000DB\u0001\u0000\u0000\u0000DE\u0001\u0000\u0000\u0000"+
		"EG\u0001\u0000\u0000\u0000FD\u0001\u0000\u0000\u0000GQ\u0005\'\u0000\u0000"+
		"HI\u0005\u0004\u0000\u0000IM\u0005&\u0000\u0000JL\u0003\u0002\u0001\u0000"+
		"KJ\u0001\u0000\u0000\u0000LO\u0001\u0000\u0000\u0000MK\u0001\u0000\u0000"+
		"\u0000MN\u0001\u0000\u0000\u0000NP\u0001\u0000\u0000\u0000OM\u0001\u0000"+
		"\u0000\u0000PR\u0005\'\u0000\u0000QH\u0001\u0000\u0000\u0000QR\u0001\u0000"+
		"\u0000\u0000R\u000b\u0001\u0000\u0000\u0000ST\u0005\u0005\u0000\u0000"+
		"TU\u0005$\u0000\u0000UV\u0005\u0013\u0000\u0000VW\u0005!\u0000\u0000W"+
		"X\u0003\u0010\b\u0000XY\u0005)\u0000\u0000YZ\u0003\u0010\b\u0000Z[\u0005"+
		"%\u0000\u0000[_\u0005&\u0000\u0000\\^\u0003\u0002\u0001\u0000]\\\u0001"+
		"\u0000\u0000\u0000^a\u0001\u0000\u0000\u0000_]\u0001\u0000\u0000\u0000"+
		"_`\u0001\u0000\u0000\u0000`b\u0001\u0000\u0000\u0000a_\u0001\u0000\u0000"+
		"\u0000bc\u0005\'\u0000\u0000c\r\u0001\u0000\u0000\u0000de\u0005\u0006"+
		"\u0000\u0000ef\u0005$\u0000\u0000fg\u0003\u0010\b\u0000gh\u0005%\u0000"+
		"\u0000hl\u0005&\u0000\u0000ik\u0003\u0002\u0001\u0000ji\u0001\u0000\u0000"+
		"\u0000kn\u0001\u0000\u0000\u0000lj\u0001\u0000\u0000\u0000lm\u0001\u0000"+
		"\u0000\u0000mo\u0001\u0000\u0000\u0000nl\u0001\u0000\u0000\u0000op\u0005"+
		"\'\u0000\u0000p\u000f\u0001\u0000\u0000\u0000qr\u0006\b\uffff\uffff\u0000"+
		"r\u008e\u0003\u0016\u000b\u0000st\u0005$\u0000\u0000tu\u0003\u0010\b\u0000"+
		"uv\u0005%\u0000\u0000v\u008e\u0001\u0000\u0000\u0000wx\u0005&\u0000\u0000"+
		"xy\u0003\u0010\b\u0000yz\u0005\'\u0000\u0000z\u008e\u0001\u0000\u0000"+
		"\u0000{|\u0007\u0000\u0000\u0000|\u008e\u0003\u0010\b\f}\u008e\u0005\u0013"+
		"\u0000\u0000~\u008e\u0003\u0018\f\u0000\u007f\u0080\u0005\u0013\u0000"+
		"\u0000\u0080\u0081\u0005(\u0000\u0000\u0081\u008e\u0005\u0013\u0000\u0000"+
		"\u0082\u0083\u0005\u0013\u0000\u0000\u0083\u0084\u0005(\u0000\u0000\u0084"+
		"\u008e\u0003\u0010\b\u0003\u0085\u0086\u0005\u0013\u0000\u0000\u0086\u0087"+
		"\u0005!\u0000\u0000\u0087\u008e\u0003\u0010\b\u0002\u0088\u0089\u0005"+
		"\u0012\u0000\u0000\u0089\u008a\u0005$\u0000\u0000\u008a\u008b\u0003\u0010"+
		"\b\u0000\u008b\u008c\u0005%\u0000\u0000\u008c\u008e\u0001\u0000\u0000"+
		"\u0000\u008dq\u0001\u0000\u0000\u0000\u008ds\u0001\u0000\u0000\u0000\u008d"+
		"w\u0001\u0000\u0000\u0000\u008d{\u0001\u0000\u0000\u0000\u008d}\u0001"+
		"\u0000\u0000\u0000\u008d~\u0001\u0000\u0000\u0000\u008d\u007f\u0001\u0000"+
		"\u0000\u0000\u008d\u0082\u0001\u0000\u0000\u0000\u008d\u0085\u0001\u0000"+
		"\u0000\u0000\u008d\u0088\u0001\u0000\u0000\u0000\u008e\u00a0\u0001\u0000"+
		"\u0000\u0000\u008f\u0090\n\u000b\u0000\u0000\u0090\u0091\u0007\u0001\u0000"+
		"\u0000\u0091\u009f\u0003\u0010\b\f\u0092\u0093\n\n\u0000\u0000\u0093\u0094"+
		"\u0007\u0002\u0000\u0000\u0094\u009f\u0003\u0010\b\u000b\u0095\u0096\n"+
		"\t\u0000\u0000\u0096\u0097\u0007\u0003\u0000\u0000\u0097\u009f\u0003\u0010"+
		"\b\n\u0098\u0099\n\b\u0000\u0000\u0099\u009a\u0007\u0004\u0000\u0000\u009a"+
		"\u009f\u0003\u0010\b\t\u009b\u009c\n\u0007\u0000\u0000\u009c\u009d\u0005"+
		"\u001a\u0000\u0000\u009d\u009f\u0003\u0010\b\b\u009e\u008f\u0001\u0000"+
		"\u0000\u0000\u009e\u0092\u0001\u0000\u0000\u0000\u009e\u0095\u0001\u0000"+
		"\u0000\u0000\u009e\u0098\u0001\u0000\u0000\u0000\u009e\u009b\u0001\u0000"+
		"\u0000\u0000\u009f\u00a2\u0001\u0000\u0000\u0000\u00a0\u009e\u0001\u0000"+
		"\u0000\u0000\u00a0\u00a1\u0001\u0000\u0000\u0000\u00a1\u0011\u0001\u0000"+
		"\u0000\u0000\u00a2\u00a0\u0001\u0000\u0000\u0000\u00a3\u00a8\u0003\u0010"+
		"\b\u0000\u00a4\u00a5\u0005)\u0000\u0000\u00a5\u00a7\u0003\u0010\b\u0000"+
		"\u00a6\u00a4\u0001\u0000\u0000\u0000\u00a7\u00aa\u0001\u0000\u0000\u0000"+
		"\u00a8\u00a6\u0001\u0000\u0000\u0000\u00a8\u00a9\u0001\u0000\u0000\u0000"+
		"\u00a9\u0013\u0001\u0000\u0000\u0000\u00aa\u00a8\u0001\u0000\u0000\u0000"+
		"\u00ab\u00ac\u0003\u0016\u000b\u0000\u00ac\u0015\u0001\u0000\u0000\u0000"+
		"\u00ad\u00b3\u0005\f\u0000\u0000\u00ae\u00b3\u0005\r\u0000\u0000\u00af"+
		"\u00b3\u0005\u000e\u0000\u0000\u00b0\u00b3\u0005\u000b\u0000\u0000\u00b1"+
		"\u00b3\u0005\u000f\u0000\u0000\u00b2\u00ad\u0001\u0000\u0000\u0000\u00b2"+
		"\u00ae\u0001\u0000\u0000\u0000\u00b2\u00af\u0001\u0000\u0000\u0000\u00b2"+
		"\u00b0\u0001\u0000\u0000\u0000\u00b2\u00b1\u0001\u0000\u0000\u0000\u00b3"+
		"\u0017\u0001\u0000\u0000\u0000\u00b4\u00b5\u0005\u0013\u0000\u0000\u00b5"+
		"\u00b9\u0005\"\u0000\u0000\u00b6\u00b7\u0005\u0013\u0000\u0000\u00b7\u00b9"+
		"\u0005#\u0000\u0000\u00b8\u00b4\u0001\u0000\u0000\u0000\u00b8\u00b6\u0001"+
		"\u0000\u0000\u0000\u00b9\u0019\u0001\u0000\u0000\u0000\u0011\u001d#\'"+
		"-5:DMQ_l\u008d\u009e\u00a0\u00a8\u00b2\u00b8";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}