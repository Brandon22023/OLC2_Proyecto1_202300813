// Generated from /home/vboxuser/Documents/OLC2_Proyecto1_202300813/Backend/parser/V4LangGrammar.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class V4LangGrammarParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, T__1=2, T__2=3, T__3=4, T__4=5, T__5=6, T__6=7, T__7=8, T__8=9, 
		T__9=10, T__10=11, T__11=12, T__12=13, T__13=14, T__14=15, T__15=16, T__16=17, 
		T__17=18, T__18=19, T__19=20, T__20=21, T__21=22, T__22=23, T__23=24, 
		T__24=25, T__25=26, T__26=27, T__27=28, T__28=29, T__29=30, T__30=31, 
		T__31=32, T__32=33, T__33=34, T__34=35, MUT=36, PRINT=37, IF=38, ELSE=39, 
		APPEND=40, LEN=41, SLICE_INDEX=42, JOIN=43, ATOI=44, PARSEFLOAT=45, TYPEOF=46, 
		FOR=47, RANGE=48, BREAK=49, NIL=50, FUNC=51, STRUCT=52, L_CURLY=53, R_CURLY=54, 
		L_PAREN=55, R_PAREN=56, L_BRACKET=57, R_BRACKET=58, DECLARE_ASSIGN=59, 
		ASSIGN=60, SEMICOLON=61, INT=62, BOOL=63, FLOAT=64, STRING=65, RUNE=66, 
		WS=67, ID=68, COMMENT=69, MULTILINE_COMMENT=70;
	public static final int
		RULE_program = 0, RULE_dcl = 1, RULE_varDcl = 2, RULE_funcDcl = 3, RULE_params = 4, 
		RULE_structDcl = 5, RULE_structBody = 6, RULE_receiver = 7, RULE_sliceDcl = 8, 
		RULE_sliceValues = 9, RULE_matrixDcl = 10, RULE_matrixValues = 11, RULE_statement = 12, 
		RULE_elseIfStmt = 13, RULE_elseStmt = 14, RULE_switchCase = 15, RULE_defaultCase = 16, 
		RULE_args = 17, RULE_expr = 18, RULE_call = 19, RULE_instanceValues = 20, 
		RULE_type = 21;
	private static String[] makeRuleNames() {
		return new String[] {
			"program", "dcl", "varDcl", "funcDcl", "params", "structDcl", "structBody", 
			"receiver", "sliceDcl", "sliceValues", "matrixDcl", "matrixValues", "statement", 
			"elseIfStmt", "elseStmt", "switchCase", "defaultCase", "args", "expr", 
			"call", "instanceValues", "type"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "','", "'type'", "'switch'", "'continue'", "'return'", "'case'", 
			"':'", "'default'", "'-'", "'!'", "'*'", "'/'", "'+'", "'<'", "'>'", 
			"'<='", "'>='", "'%'", "'=='", "'!='", "'+='", "'-='", "'*='", "'/='", 
			"'&&'", "'||'", "'new'", "'++'", "'--'", "'.'", "'int'", "'bool'", "'float64'", 
			"'string'", "'rune'", "'mut'", "'fmt.Println'", "'if'", "'else'", "'append'", 
			"'len'", "'slices.Index'", "'strings.Join'", "'strconv.Atoi'", "'strconv.ParseFloat'", 
			"'reflect.TypeOf'", "'for'", "'range'", "'break'", "'nil'", "'func'", 
			"'struct'", "'{'", "'}'", "'('", "')'", "'['", "']'", "':='", "'='", 
			"';'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			null, null, null, null, null, null, null, null, null, null, null, null, 
			"MUT", "PRINT", "IF", "ELSE", "APPEND", "LEN", "SLICE_INDEX", "JOIN", 
			"ATOI", "PARSEFLOAT", "TYPEOF", "FOR", "RANGE", "BREAK", "NIL", "FUNC", 
			"STRUCT", "L_CURLY", "R_CURLY", "L_PAREN", "R_PAREN", "L_BRACKET", "R_BRACKET", 
			"DECLARE_ASSIGN", "ASSIGN", "SEMICOLON", "INT", "BOOL", "FLOAT", "STRING", 
			"RUNE", "WS", "ID", "COMMENT", "MULTILINE_COMMENT"
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
	public String getGrammarFileName() { return "V4LangGrammar.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public V4LangGrammarParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ProgramContext extends ParserRuleContext {
		public List<DclContext> dcl() {
			return getRuleContexts(DclContext.class);
		}
		public DclContext dcl(int i) {
			return getRuleContext(DclContext.class,i);
		}
		public ProgramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_program; }
	}

	public final ProgramContext program() throws RecognitionException {
		ProgramContext _localctx = new ProgramContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_program);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(47);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & -4418313327768238532L) != 0) || ((((_la - 64)) & ~0x3f) == 0 && ((1L << (_la - 64)) & 23L) != 0)) {
				{
				{
				setState(44);
				dcl();
				}
				}
				setState(49);
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
	public static class DclContext extends ParserRuleContext {
		public VarDclContext varDcl() {
			return getRuleContext(VarDclContext.class,0);
		}
		public TerminalNode SEMICOLON() { return getToken(V4LangGrammarParser.SEMICOLON, 0); }
		public MatrixDclContext matrixDcl() {
			return getRuleContext(MatrixDclContext.class,0);
		}
		public SliceDclContext sliceDcl() {
			return getRuleContext(SliceDclContext.class,0);
		}
		public StatementContext statement() {
			return getRuleContext(StatementContext.class,0);
		}
		public FuncDclContext funcDcl() {
			return getRuleContext(FuncDclContext.class,0);
		}
		public StructDclContext structDcl() {
			return getRuleContext(StructDclContext.class,0);
		}
		public DclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_dcl; }
	}

	public final DclContext dcl() throws RecognitionException {
		DclContext _localctx = new DclContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_dcl);
		int _la;
		try {
			setState(74);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(50);
				varDcl();
				setState(52);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMICOLON) {
					{
					setState(51);
					match(SEMICOLON);
					}
				}

				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(54);
				matrixDcl();
				setState(56);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMICOLON) {
					{
					setState(55);
					match(SEMICOLON);
					}
				}

				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(58);
				sliceDcl();
				setState(60);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMICOLON) {
					{
					setState(59);
					match(SEMICOLON);
					}
				}

				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(62);
				statement();
				setState(64);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMICOLON) {
					{
					setState(63);
					match(SEMICOLON);
					}
				}

				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(66);
				funcDcl();
				setState(68);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMICOLON) {
					{
					setState(67);
					match(SEMICOLON);
					}
				}

				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(70);
				structDcl();
				setState(72);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==SEMICOLON) {
					{
					setState(71);
					match(SEMICOLON);
					}
				}

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
		public TerminalNode MUT() { return getToken(V4LangGrammarParser.MUT, 0); }
		public List<TerminalNode> ID() { return getTokens(V4LangGrammarParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(V4LangGrammarParser.ID, i);
		}
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public TerminalNode ASSIGN() { return getToken(V4LangGrammarParser.ASSIGN, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode L_CURLY() { return getToken(V4LangGrammarParser.L_CURLY, 0); }
		public TerminalNode R_CURLY() { return getToken(V4LangGrammarParser.R_CURLY, 0); }
		public TerminalNode DECLARE_ASSIGN() { return getToken(V4LangGrammarParser.DECLARE_ASSIGN, 0); }
		public VarDclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_varDcl; }
	}

	public final VarDclContext varDcl() throws RecognitionException {
		VarDclContext _localctx = new VarDclContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_varDcl);
		try {
			setState(97);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(76);
				match(MUT);
				setState(77);
				match(ID);
				setState(78);
				type();
				setState(79);
				match(ASSIGN);
				setState(80);
				expr(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(82);
				match(MUT);
				setState(83);
				match(ID);
				setState(84);
				type();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(85);
				match(MUT);
				setState(86);
				match(ID);
				setState(87);
				match(L_CURLY);
				setState(88);
				match(R_CURLY);
				setState(89);
				type();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(90);
				match(ID);
				setState(91);
				match(DECLARE_ASSIGN);
				setState(92);
				expr(0);
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(93);
				match(ID);
				setState(94);
				type();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(95);
				match(ID);
				setState(96);
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
	public static class FuncDclContext extends ParserRuleContext {
		public TerminalNode FUNC() { return getToken(V4LangGrammarParser.FUNC, 0); }
		public TerminalNode ID() { return getToken(V4LangGrammarParser.ID, 0); }
		public TerminalNode L_PAREN() { return getToken(V4LangGrammarParser.L_PAREN, 0); }
		public TerminalNode R_PAREN() { return getToken(V4LangGrammarParser.R_PAREN, 0); }
		public TerminalNode L_CURLY() { return getToken(V4LangGrammarParser.L_CURLY, 0); }
		public TerminalNode R_CURLY() { return getToken(V4LangGrammarParser.R_CURLY, 0); }
		public ParamsContext params() {
			return getRuleContext(ParamsContext.class,0);
		}
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public List<DclContext> dcl() {
			return getRuleContexts(DclContext.class);
		}
		public DclContext dcl(int i) {
			return getRuleContext(DclContext.class,i);
		}
		public ReceiverContext receiver() {
			return getRuleContext(ReceiverContext.class,0);
		}
		public FuncDclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funcDcl; }
	}

	public final FuncDclContext funcDcl() throws RecognitionException {
		FuncDclContext _localctx = new FuncDclContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_funcDcl);
		int _la;
		try {
			setState(137);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,15,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(99);
				match(FUNC);
				setState(100);
				match(ID);
				setState(101);
				match(L_PAREN);
				setState(103);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==ID) {
					{
					setState(102);
					params();
					}
				}

				setState(105);
				match(R_PAREN);
				setState(107);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 66571993088L) != 0)) {
					{
					setState(106);
					type();
					}
				}

				setState(109);
				match(L_CURLY);
				setState(113);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while ((((_la) & ~0x3f) == 0 && ((1L << _la) & -4418313327768238532L) != 0) || ((((_la - 64)) & ~0x3f) == 0 && ((1L << (_la - 64)) & 23L) != 0)) {
					{
					{
					setState(110);
					dcl();
					}
					}
					setState(115);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(116);
				match(R_CURLY);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(117);
				match(FUNC);
				setState(118);
				receiver();
				setState(119);
				match(ID);
				setState(120);
				match(L_PAREN);
				setState(122);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==ID) {
					{
					setState(121);
					params();
					}
				}

				setState(124);
				match(R_PAREN);
				setState(126);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 66571993088L) != 0)) {
					{
					setState(125);
					type();
					}
				}

				setState(128);
				match(L_CURLY);
				setState(132);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while ((((_la) & ~0x3f) == 0 && ((1L << _la) & -4418313327768238532L) != 0) || ((((_la - 64)) & ~0x3f) == 0 && ((1L << (_la - 64)) & 23L) != 0)) {
					{
					{
					setState(129);
					dcl();
					}
					}
					setState(134);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(135);
				match(R_CURLY);
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
	public static class ParamsContext extends ParserRuleContext {
		public List<TerminalNode> ID() { return getTokens(V4LangGrammarParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(V4LangGrammarParser.ID, i);
		}
		public List<TypeContext> type() {
			return getRuleContexts(TypeContext.class);
		}
		public TypeContext type(int i) {
			return getRuleContext(TypeContext.class,i);
		}
		public List<TerminalNode> L_BRACKET() { return getTokens(V4LangGrammarParser.L_BRACKET); }
		public TerminalNode L_BRACKET(int i) {
			return getToken(V4LangGrammarParser.L_BRACKET, i);
		}
		public List<TerminalNode> R_BRACKET() { return getTokens(V4LangGrammarParser.R_BRACKET); }
		public TerminalNode R_BRACKET(int i) {
			return getToken(V4LangGrammarParser.R_BRACKET, i);
		}
		public ParamsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_params; }
	}

	public final ParamsContext params() throws RecognitionException {
		ParamsContext _localctx = new ParamsContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_params);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(139);
			match(ID);
			setState(142);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==L_BRACKET) {
				{
				setState(140);
				match(L_BRACKET);
				setState(141);
				match(R_BRACKET);
				}
			}

			setState(144);
			type();
			setState(154);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__0) {
				{
				{
				setState(145);
				match(T__0);
				setState(146);
				match(ID);
				setState(149);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==L_BRACKET) {
					{
					setState(147);
					match(L_BRACKET);
					setState(148);
					match(R_BRACKET);
					}
				}

				setState(151);
				type();
				}
				}
				setState(156);
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
		public TerminalNode ID() { return getToken(V4LangGrammarParser.ID, 0); }
		public TerminalNode STRUCT() { return getToken(V4LangGrammarParser.STRUCT, 0); }
		public TerminalNode L_CURLY() { return getToken(V4LangGrammarParser.L_CURLY, 0); }
		public TerminalNode R_CURLY() { return getToken(V4LangGrammarParser.R_CURLY, 0); }
		public List<StructBodyContext> structBody() {
			return getRuleContexts(StructBodyContext.class);
		}
		public StructBodyContext structBody(int i) {
			return getRuleContext(StructBodyContext.class,i);
		}
		public StructDclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_structDcl; }
	}

	public final StructDclContext structDcl() throws RecognitionException {
		StructDclContext _localctx = new StructDclContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_structDcl);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(157);
			match(T__1);
			setState(158);
			match(ID);
			setState(159);
			match(STRUCT);
			setState(160);
			match(L_CURLY);
			setState(164);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (((((_la - 36)) & ~0x3f) == 0 && ((1L << (_la - 36)) & 4295000065L) != 0)) {
				{
				{
				setState(161);
				structBody();
				}
				}
				setState(166);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(167);
			match(R_CURLY);
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
	public static class StructBodyContext extends ParserRuleContext {
		public VarDclContext varDcl() {
			return getRuleContext(VarDclContext.class,0);
		}
		public FuncDclContext funcDcl() {
			return getRuleContext(FuncDclContext.class,0);
		}
		public StructBodyContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_structBody; }
	}

	public final StructBodyContext structBody() throws RecognitionException {
		StructBodyContext _localctx = new StructBodyContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_structBody);
		try {
			setState(171);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case MUT:
			case ID:
				enterOuterAlt(_localctx, 1);
				{
				setState(169);
				varDcl();
				}
				break;
			case FUNC:
				enterOuterAlt(_localctx, 2);
				{
				setState(170);
				funcDcl();
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
	public static class ReceiverContext extends ParserRuleContext {
		public Token this_;
		public Token structName;
		public TerminalNode L_PAREN() { return getToken(V4LangGrammarParser.L_PAREN, 0); }
		public TerminalNode R_PAREN() { return getToken(V4LangGrammarParser.R_PAREN, 0); }
		public List<TerminalNode> ID() { return getTokens(V4LangGrammarParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(V4LangGrammarParser.ID, i);
		}
		public ReceiverContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_receiver; }
	}

	public final ReceiverContext receiver() throws RecognitionException {
		ReceiverContext _localctx = new ReceiverContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_receiver);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(173);
			match(L_PAREN);
			setState(174);
			((ReceiverContext)_localctx).this_ = match(ID);
			setState(175);
			((ReceiverContext)_localctx).structName = match(ID);
			setState(176);
			match(R_PAREN);
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
	public static class SliceDclContext extends ParserRuleContext {
		public TerminalNode MUT() { return getToken(V4LangGrammarParser.MUT, 0); }
		public TerminalNode ID() { return getToken(V4LangGrammarParser.ID, 0); }
		public TerminalNode DECLARE_ASSIGN() { return getToken(V4LangGrammarParser.DECLARE_ASSIGN, 0); }
		public TerminalNode L_BRACKET() { return getToken(V4LangGrammarParser.L_BRACKET, 0); }
		public TerminalNode R_BRACKET() { return getToken(V4LangGrammarParser.R_BRACKET, 0); }
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public TerminalNode L_CURLY() { return getToken(V4LangGrammarParser.L_CURLY, 0); }
		public SliceValuesContext sliceValues() {
			return getRuleContext(SliceValuesContext.class,0);
		}
		public TerminalNode R_CURLY() { return getToken(V4LangGrammarParser.R_CURLY, 0); }
		public SliceDclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sliceDcl; }
	}

	public final SliceDclContext sliceDcl() throws RecognitionException {
		SliceDclContext _localctx = new SliceDclContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_sliceDcl);
		try {
			setState(190);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case MUT:
				enterOuterAlt(_localctx, 1);
				{
				setState(178);
				match(MUT);
				setState(179);
				match(ID);
				setState(180);
				match(DECLARE_ASSIGN);
				setState(181);
				match(L_BRACKET);
				setState(182);
				match(R_BRACKET);
				setState(183);
				type();
				setState(184);
				match(L_CURLY);
				setState(185);
				sliceValues();
				setState(186);
				match(R_CURLY);
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 2);
				{
				setState(188);
				match(ID);
				setState(189);
				type();
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
	public static class SliceValuesContext extends ParserRuleContext {
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public SliceValuesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sliceValues; }
	}

	public final SliceValuesContext sliceValues() throws RecognitionException {
		SliceValuesContext _localctx = new SliceValuesContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_sliceValues);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(192);
			expr(0);
			setState(197);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__0) {
				{
				{
				setState(193);
				match(T__0);
				setState(194);
				expr(0);
				}
				}
				setState(199);
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
	public static class MatrixDclContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(V4LangGrammarParser.ID, 0); }
		public TerminalNode DECLARE_ASSIGN() { return getToken(V4LangGrammarParser.DECLARE_ASSIGN, 0); }
		public List<TerminalNode> L_BRACKET() { return getTokens(V4LangGrammarParser.L_BRACKET); }
		public TerminalNode L_BRACKET(int i) {
			return getToken(V4LangGrammarParser.L_BRACKET, i);
		}
		public List<TerminalNode> R_BRACKET() { return getTokens(V4LangGrammarParser.R_BRACKET); }
		public TerminalNode R_BRACKET(int i) {
			return getToken(V4LangGrammarParser.R_BRACKET, i);
		}
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public TerminalNode L_CURLY() { return getToken(V4LangGrammarParser.L_CURLY, 0); }
		public MatrixValuesContext matrixValues() {
			return getRuleContext(MatrixValuesContext.class,0);
		}
		public TerminalNode R_CURLY() { return getToken(V4LangGrammarParser.R_CURLY, 0); }
		public MatrixDclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_matrixDcl; }
	}

	public final MatrixDclContext matrixDcl() throws RecognitionException {
		MatrixDclContext _localctx = new MatrixDclContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_matrixDcl);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(200);
			match(ID);
			setState(201);
			match(DECLARE_ASSIGN);
			setState(202);
			match(L_BRACKET);
			setState(203);
			match(R_BRACKET);
			setState(204);
			match(L_BRACKET);
			setState(205);
			match(R_BRACKET);
			setState(206);
			type();
			setState(207);
			match(L_CURLY);
			setState(208);
			matrixValues();
			setState(209);
			match(R_CURLY);
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
	public static class MatrixValuesContext extends ParserRuleContext {
		public List<TerminalNode> L_CURLY() { return getTokens(V4LangGrammarParser.L_CURLY); }
		public TerminalNode L_CURLY(int i) {
			return getToken(V4LangGrammarParser.L_CURLY, i);
		}
		public List<SliceValuesContext> sliceValues() {
			return getRuleContexts(SliceValuesContext.class);
		}
		public SliceValuesContext sliceValues(int i) {
			return getRuleContext(SliceValuesContext.class,i);
		}
		public List<TerminalNode> R_CURLY() { return getTokens(V4LangGrammarParser.R_CURLY); }
		public TerminalNode R_CURLY(int i) {
			return getToken(V4LangGrammarParser.R_CURLY, i);
		}
		public MatrixValuesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_matrixValues; }
	}

	public final MatrixValuesContext matrixValues() throws RecognitionException {
		MatrixValuesContext _localctx = new MatrixValuesContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_matrixValues);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(211);
			match(L_CURLY);
			setState(212);
			sliceValues();
			setState(213);
			match(R_CURLY);
			setState(221);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,23,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(214);
					match(T__0);
					setState(215);
					match(L_CURLY);
					setState(216);
					sliceValues();
					setState(217);
					match(R_CURLY);
					}
					} 
				}
				setState(223);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,23,_ctx);
			}
			setState(225);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__0) {
				{
				setState(224);
				match(T__0);
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
	public static class StatementContext extends ParserRuleContext {
		public StatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_statement; }
	 
		public StatementContext() { }
		public void copyFrom(StatementContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ContinueStmtContext extends StatementContext {
		public ContinueStmtContext(StatementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SwitchStmtContext extends StatementContext {
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode L_CURLY() { return getToken(V4LangGrammarParser.L_CURLY, 0); }
		public TerminalNode R_CURLY() { return getToken(V4LangGrammarParser.R_CURLY, 0); }
		public List<SwitchCaseContext> switchCase() {
			return getRuleContexts(SwitchCaseContext.class);
		}
		public SwitchCaseContext switchCase(int i) {
			return getRuleContext(SwitchCaseContext.class,i);
		}
		public DefaultCaseContext defaultCase() {
			return getRuleContext(DefaultCaseContext.class,0);
		}
		public SwitchStmtContext(StatementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class PrintStmtContext extends StatementContext {
		public TerminalNode PRINT() { return getToken(V4LangGrammarParser.PRINT, 0); }
		public TerminalNode L_PAREN() { return getToken(V4LangGrammarParser.L_PAREN, 0); }
		public TerminalNode R_PAREN() { return getToken(V4LangGrammarParser.R_PAREN, 0); }
		public ArgsContext args() {
			return getRuleContext(ArgsContext.class,0);
		}
		public PrintStmtContext(StatementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IfStmtContext extends StatementContext {
		public TerminalNode IF() { return getToken(V4LangGrammarParser.IF, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public StatementContext statement() {
			return getRuleContext(StatementContext.class,0);
		}
		public List<ElseIfStmtContext> elseIfStmt() {
			return getRuleContexts(ElseIfStmtContext.class);
		}
		public ElseIfStmtContext elseIfStmt(int i) {
			return getRuleContext(ElseIfStmtContext.class,i);
		}
		public ElseStmtContext elseStmt() {
			return getRuleContext(ElseStmtContext.class,0);
		}
		public IfStmtContext(StatementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExprStmtContext extends StatementContext {
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public ExprStmtContext(StatementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class BreakStmtContext extends StatementContext {
		public TerminalNode BREAK() { return getToken(V4LangGrammarParser.BREAK, 0); }
		public BreakStmtContext(StatementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class BlockStmtContext extends StatementContext {
		public TerminalNode L_CURLY() { return getToken(V4LangGrammarParser.L_CURLY, 0); }
		public TerminalNode R_CURLY() { return getToken(V4LangGrammarParser.R_CURLY, 0); }
		public List<DclContext> dcl() {
			return getRuleContexts(DclContext.class);
		}
		public DclContext dcl(int i) {
			return getRuleContext(DclContext.class,i);
		}
		public BlockStmtContext(StatementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ForRangeStmtContext extends StatementContext {
		public TerminalNode FOR() { return getToken(V4LangGrammarParser.FOR, 0); }
		public List<TerminalNode> ID() { return getTokens(V4LangGrammarParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(V4LangGrammarParser.ID, i);
		}
		public TerminalNode DECLARE_ASSIGN() { return getToken(V4LangGrammarParser.DECLARE_ASSIGN, 0); }
		public TerminalNode RANGE() { return getToken(V4LangGrammarParser.RANGE, 0); }
		public StatementContext statement() {
			return getRuleContext(StatementContext.class,0);
		}
		public ForRangeStmtContext(StatementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ForIncrementStmtContext extends StatementContext {
		public TerminalNode FOR() { return getToken(V4LangGrammarParser.FOR, 0); }
		public VarDclContext varDcl() {
			return getRuleContext(VarDclContext.class,0);
		}
		public List<TerminalNode> SEMICOLON() { return getTokens(V4LangGrammarParser.SEMICOLON); }
		public TerminalNode SEMICOLON(int i) {
			return getToken(V4LangGrammarParser.SEMICOLON, i);
		}
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public StatementContext statement() {
			return getRuleContext(StatementContext.class,0);
		}
		public ForIncrementStmtContext(StatementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ForStmtContext extends StatementContext {
		public TerminalNode FOR() { return getToken(V4LangGrammarParser.FOR, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public StatementContext statement() {
			return getRuleContext(StatementContext.class,0);
		}
		public ForStmtContext(StatementContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ReturnStmtContext extends StatementContext {
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public ReturnStmtContext(StatementContext ctx) { copyFrom(ctx); }
	}

	public final StatementContext statement() throws RecognitionException {
		StatementContext _localctx = new StatementContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_statement);
		int _la;
		try {
			int _alt;
			setState(294);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,32,_ctx) ) {
			case 1:
				_localctx = new ExprStmtContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(227);
				expr(0);
				}
				break;
			case 2:
				_localctx = new PrintStmtContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(228);
				match(PRINT);
				setState(229);
				match(L_PAREN);
				setState(231);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (((((_la - 9)) & ~0x3f) == 0 && ((1L << (_la - 9)) & 856055836861267971L) != 0)) {
					{
					setState(230);
					args();
					}
				}

				setState(233);
				match(R_PAREN);
				}
				break;
			case 3:
				_localctx = new BlockStmtContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(234);
				match(L_CURLY);
				setState(238);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while ((((_la) & ~0x3f) == 0 && ((1L << _la) & -4418313327768238532L) != 0) || ((((_la - 64)) & ~0x3f) == 0 && ((1L << (_la - 64)) & 23L) != 0)) {
					{
					{
					setState(235);
					dcl();
					}
					}
					setState(240);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(241);
				match(R_CURLY);
				}
				break;
			case 4:
				_localctx = new IfStmtContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(242);
				match(IF);
				setState(243);
				expr(0);
				setState(244);
				statement();
				setState(248);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,27,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(245);
						elseIfStmt();
						}
						} 
					}
					setState(250);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,27,_ctx);
				}
				setState(252);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,28,_ctx) ) {
				case 1:
					{
					setState(251);
					elseStmt();
					}
					break;
				}
				}
				break;
			case 5:
				_localctx = new SwitchStmtContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(254);
				match(T__2);
				setState(255);
				expr(0);
				setState(256);
				match(L_CURLY);
				setState(260);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==T__5) {
					{
					{
					setState(257);
					switchCase();
					}
					}
					setState(262);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(264);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==T__7) {
					{
					setState(263);
					defaultCase();
					}
				}

				setState(266);
				match(R_CURLY);
				}
				break;
			case 6:
				_localctx = new ForIncrementStmtContext(_localctx);
				enterOuterAlt(_localctx, 6);
				{
				setState(268);
				match(FOR);
				setState(269);
				varDcl();
				setState(270);
				match(SEMICOLON);
				setState(271);
				expr(0);
				setState(272);
				match(SEMICOLON);
				setState(273);
				expr(0);
				setState(274);
				statement();
				}
				break;
			case 7:
				_localctx = new ForStmtContext(_localctx);
				enterOuterAlt(_localctx, 7);
				{
				setState(276);
				match(FOR);
				setState(277);
				expr(0);
				setState(278);
				statement();
				}
				break;
			case 8:
				_localctx = new ForRangeStmtContext(_localctx);
				enterOuterAlt(_localctx, 8);
				{
				setState(280);
				match(FOR);
				setState(281);
				match(ID);
				setState(282);
				match(T__0);
				setState(283);
				match(ID);
				setState(284);
				match(DECLARE_ASSIGN);
				setState(285);
				match(RANGE);
				setState(286);
				match(ID);
				setState(287);
				statement();
				}
				break;
			case 9:
				_localctx = new BreakStmtContext(_localctx);
				enterOuterAlt(_localctx, 9);
				{
				setState(288);
				match(BREAK);
				}
				break;
			case 10:
				_localctx = new ContinueStmtContext(_localctx);
				enterOuterAlt(_localctx, 10);
				{
				setState(289);
				match(T__3);
				}
				break;
			case 11:
				_localctx = new ReturnStmtContext(_localctx);
				enterOuterAlt(_localctx, 11);
				{
				setState(290);
				match(T__4);
				setState(292);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,31,_ctx) ) {
				case 1:
					{
					setState(291);
					expr(0);
					}
					break;
				}
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
	public static class ElseIfStmtContext extends ParserRuleContext {
		public TerminalNode ELSE() { return getToken(V4LangGrammarParser.ELSE, 0); }
		public TerminalNode IF() { return getToken(V4LangGrammarParser.IF, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public StatementContext statement() {
			return getRuleContext(StatementContext.class,0);
		}
		public ElseIfStmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_elseIfStmt; }
	}

	public final ElseIfStmtContext elseIfStmt() throws RecognitionException {
		ElseIfStmtContext _localctx = new ElseIfStmtContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_elseIfStmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(296);
			match(ELSE);
			setState(297);
			match(IF);
			setState(298);
			expr(0);
			setState(299);
			statement();
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
	public static class ElseStmtContext extends ParserRuleContext {
		public TerminalNode ELSE() { return getToken(V4LangGrammarParser.ELSE, 0); }
		public StatementContext statement() {
			return getRuleContext(StatementContext.class,0);
		}
		public ElseStmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_elseStmt; }
	}

	public final ElseStmtContext elseStmt() throws RecognitionException {
		ElseStmtContext _localctx = new ElseStmtContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_elseStmt);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(301);
			match(ELSE);
			setState(302);
			statement();
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
	public static class SwitchCaseContext extends ParserRuleContext {
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public List<DclContext> dcl() {
			return getRuleContexts(DclContext.class);
		}
		public DclContext dcl(int i) {
			return getRuleContext(DclContext.class,i);
		}
		public SwitchCaseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_switchCase; }
	}

	public final SwitchCaseContext switchCase() throws RecognitionException {
		SwitchCaseContext _localctx = new SwitchCaseContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_switchCase);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(304);
			match(T__5);
			setState(305);
			expr(0);
			setState(306);
			match(T__6);
			setState(310);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & -4418313327768238532L) != 0) || ((((_la - 64)) & ~0x3f) == 0 && ((1L << (_la - 64)) & 23L) != 0)) {
				{
				{
				setState(307);
				dcl();
				}
				}
				setState(312);
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
	public static class DefaultCaseContext extends ParserRuleContext {
		public List<DclContext> dcl() {
			return getRuleContexts(DclContext.class);
		}
		public DclContext dcl(int i) {
			return getRuleContext(DclContext.class,i);
		}
		public DefaultCaseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_defaultCase; }
	}

	public final DefaultCaseContext defaultCase() throws RecognitionException {
		DefaultCaseContext _localctx = new DefaultCaseContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_defaultCase);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(313);
			match(T__7);
			setState(314);
			match(T__6);
			setState(318);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & -4418313327768238532L) != 0) || ((((_la - 64)) & ~0x3f) == 0 && ((1L << (_la - 64)) & 23L) != 0)) {
				{
				{
				setState(315);
				dcl();
				}
				}
				setState(320);
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
	public static class ArgsContext extends ParserRuleContext {
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public ArgsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_args; }
	}

	public final ArgsContext args() throws RecognitionException {
		ArgsContext _localctx = new ArgsContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_args);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(321);
			expr(0);
			setState(326);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__0) {
				{
				{
				setState(322);
				match(T__0);
				setState(323);
				expr(0);
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
	public static class ExprContext extends ParserRuleContext {
		public ExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expr; }
	 
		public ExprContext() { }
		public void copyFrom(ExprContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class MulDivContext extends ExprContext {
		public Token op;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public MulDivContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class GetValueIndexContext extends ExprContext {
		public TerminalNode ID() { return getToken(V4LangGrammarParser.ID, 0); }
		public TerminalNode L_BRACKET() { return getToken(V4LangGrammarParser.L_BRACKET, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode R_BRACKET() { return getToken(V4LangGrammarParser.R_BRACKET, 0); }
		public GetValueIndexContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ParensContext extends ExprContext {
		public TerminalNode L_PAREN() { return getToken(V4LangGrammarParser.L_PAREN, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode R_PAREN() { return getToken(V4LangGrammarParser.R_PAREN, 0); }
		public ParensContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class LogicalContext extends ExprContext {
		public Token op;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public LogicalContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StringContext extends ExprContext {
		public TerminalNode STRING() { return getToken(V4LangGrammarParser.STRING, 0); }
		public StringContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IntContext extends ExprContext {
		public TerminalNode INT() { return getToken(V4LangGrammarParser.INT, 0); }
		public IntContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IdentifierContext extends ExprContext {
		public TerminalNode ID() { return getToken(V4LangGrammarParser.ID, 0); }
		public IdentifierContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SetValueIndexContext extends ExprContext {
		public ExprContext valor;
		public TerminalNode ID() { return getToken(V4LangGrammarParser.ID, 0); }
		public TerminalNode L_BRACKET() { return getToken(V4LangGrammarParser.L_BRACKET, 0); }
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode R_BRACKET() { return getToken(V4LangGrammarParser.R_BRACKET, 0); }
		public TerminalNode ASSIGN() { return getToken(V4LangGrammarParser.ASSIGN, 0); }
		public SetValueIndexContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class GetValueMatrixContext extends ExprContext {
		public TerminalNode ID() { return getToken(V4LangGrammarParser.ID, 0); }
		public List<TerminalNode> L_BRACKET() { return getTokens(V4LangGrammarParser.L_BRACKET); }
		public TerminalNode L_BRACKET(int i) {
			return getToken(V4LangGrammarParser.L_BRACKET, i);
		}
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public List<TerminalNode> R_BRACKET() { return getTokens(V4LangGrammarParser.R_BRACKET); }
		public TerminalNode R_BRACKET(int i) {
			return getToken(V4LangGrammarParser.R_BRACKET, i);
		}
		public GetValueMatrixContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SliceWithValuesContext extends ExprContext {
		public TerminalNode L_BRACKET() { return getToken(V4LangGrammarParser.L_BRACKET, 0); }
		public TerminalNode R_BRACKET() { return getToken(V4LangGrammarParser.R_BRACKET, 0); }
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public TerminalNode L_CURLY() { return getToken(V4LangGrammarParser.L_CURLY, 0); }
		public SliceValuesContext sliceValues() {
			return getRuleContext(SliceValuesContext.class,0);
		}
		public TerminalNode R_CURLY() { return getToken(V4LangGrammarParser.R_CURLY, 0); }
		public SliceWithValuesContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class EqualityContext extends ExprContext {
		public Token op;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public EqualityContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class BooleanContext extends ExprContext {
		public TerminalNode BOOL() { return getToken(V4LangGrammarParser.BOOL, 0); }
		public BooleanContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class NewInstanceContext extends ExprContext {
		public TerminalNode ID() { return getToken(V4LangGrammarParser.ID, 0); }
		public TerminalNode L_PAREN() { return getToken(V4LangGrammarParser.L_PAREN, 0); }
		public TerminalNode R_PAREN() { return getToken(V4LangGrammarParser.R_PAREN, 0); }
		public ArgsContext args() {
			return getRuleContext(ArgsContext.class,0);
		}
		public NewInstanceContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SetValueMatrixContext extends ExprContext {
		public TerminalNode ID() { return getToken(V4LangGrammarParser.ID, 0); }
		public List<TerminalNode> L_BRACKET() { return getTokens(V4LangGrammarParser.L_BRACKET); }
		public TerminalNode L_BRACKET(int i) {
			return getToken(V4LangGrammarParser.L_BRACKET, i);
		}
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public List<TerminalNode> R_BRACKET() { return getTokens(V4LangGrammarParser.R_BRACKET); }
		public TerminalNode R_BRACKET(int i) {
			return getToken(V4LangGrammarParser.R_BRACKET, i);
		}
		public TerminalNode ASSIGN() { return getToken(V4LangGrammarParser.ASSIGN, 0); }
		public SetValueMatrixContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ModContext extends ExprContext {
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public ModContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SliceIndexContext extends ExprContext {
		public TerminalNode SLICE_INDEX() { return getToken(V4LangGrammarParser.SLICE_INDEX, 0); }
		public TerminalNode L_PAREN() { return getToken(V4LangGrammarParser.L_PAREN, 0); }
		public TerminalNode ID() { return getToken(V4LangGrammarParser.ID, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode R_PAREN() { return getToken(V4LangGrammarParser.R_PAREN, 0); }
		public SliceIndexContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class TypeOfContext extends ExprContext {
		public TerminalNode TYPEOF() { return getToken(V4LangGrammarParser.TYPEOF, 0); }
		public TerminalNode L_PAREN() { return getToken(V4LangGrammarParser.L_PAREN, 0); }
		public TerminalNode ID() { return getToken(V4LangGrammarParser.ID, 0); }
		public TerminalNode R_PAREN() { return getToken(V4LangGrammarParser.R_PAREN, 0); }
		public TypeOfContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AddSubContext extends ExprContext {
		public Token op;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public AddSubContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StructInstanceContext extends ExprContext {
		public TerminalNode ID() { return getToken(V4LangGrammarParser.ID, 0); }
		public TerminalNode L_CURLY() { return getToken(V4LangGrammarParser.L_CURLY, 0); }
		public InstanceValuesContext instanceValues() {
			return getRuleContext(InstanceValuesContext.class,0);
		}
		public TerminalNode R_CURLY() { return getToken(V4LangGrammarParser.R_CURLY, 0); }
		public StructInstanceContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class RelationalContext extends ExprContext {
		public Token op;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public RelationalContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class CallStmtContext extends ExprContext {
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public List<CallContext> call() {
			return getRuleContexts(CallContext.class);
		}
		public CallContext call(int i) {
			return getRuleContext(CallContext.class,i);
		}
		public CallStmtContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class NilContext extends ExprContext {
		public TerminalNode NIL() { return getToken(V4LangGrammarParser.NIL, 0); }
		public NilContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StringJoinContext extends ExprContext {
		public TerminalNode JOIN() { return getToken(V4LangGrammarParser.JOIN, 0); }
		public TerminalNode L_PAREN() { return getToken(V4LangGrammarParser.L_PAREN, 0); }
		public TerminalNode ID() { return getToken(V4LangGrammarParser.ID, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode R_PAREN() { return getToken(V4LangGrammarParser.R_PAREN, 0); }
		public StringJoinContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FloatContext extends ExprContext {
		public TerminalNode FLOAT() { return getToken(V4LangGrammarParser.FLOAT, 0); }
		public FloatContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class NotContext extends ExprContext {
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public NotContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AtoiContext extends ExprContext {
		public TerminalNode ATOI() { return getToken(V4LangGrammarParser.ATOI, 0); }
		public TerminalNode L_PAREN() { return getToken(V4LangGrammarParser.L_PAREN, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode R_PAREN() { return getToken(V4LangGrammarParser.R_PAREN, 0); }
		public AtoiContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ParseFloatContext extends ExprContext {
		public TerminalNode PARSEFLOAT() { return getToken(V4LangGrammarParser.PARSEFLOAT, 0); }
		public TerminalNode L_PAREN() { return getToken(V4LangGrammarParser.L_PAREN, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode R_PAREN() { return getToken(V4LangGrammarParser.R_PAREN, 0); }
		public ParseFloatContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AppendContext extends ExprContext {
		public TerminalNode APPEND() { return getToken(V4LangGrammarParser.APPEND, 0); }
		public TerminalNode L_PAREN() { return getToken(V4LangGrammarParser.L_PAREN, 0); }
		public TerminalNode ID() { return getToken(V4LangGrammarParser.ID, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode R_PAREN() { return getToken(V4LangGrammarParser.R_PAREN, 0); }
		public AppendContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class LenContext extends ExprContext {
		public TerminalNode LEN() { return getToken(V4LangGrammarParser.LEN, 0); }
		public TerminalNode L_PAREN() { return getToken(V4LangGrammarParser.L_PAREN, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode R_PAREN() { return getToken(V4LangGrammarParser.R_PAREN, 0); }
		public LenContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AssignArithmeticContext extends ExprContext {
		public Token op;
		public TerminalNode ID() { return getToken(V4LangGrammarParser.ID, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public AssignArithmeticContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SliceContext extends ExprContext {
		public TerminalNode L_CURLY() { return getToken(V4LangGrammarParser.L_CURLY, 0); }
		public TerminalNode R_CURLY() { return getToken(V4LangGrammarParser.R_CURLY, 0); }
		public List<SliceValuesContext> sliceValues() {
			return getRuleContexts(SliceValuesContext.class);
		}
		public SliceValuesContext sliceValues(int i) {
			return getRuleContext(SliceValuesContext.class,i);
		}
		public SliceContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AssignContext extends ExprContext {
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode ASSIGN() { return getToken(V4LangGrammarParser.ASSIGN, 0); }
		public AssignContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class NegateContext extends ExprContext {
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public NegateContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IncrementDecrementContext extends ExprContext {
		public Token op;
		public TerminalNode ID() { return getToken(V4LangGrammarParser.ID, 0); }
		public IncrementDecrementContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class RuneContext extends ExprContext {
		public TerminalNode RUNE() { return getToken(V4LangGrammarParser.RUNE, 0); }
		public RuneContext(ExprContext ctx) { copyFrom(ctx); }
	}

	public final ExprContext expr() throws RecognitionException {
		return expr(0);
	}

	private ExprContext expr(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExprContext _localctx = new ExprContext(_ctx, _parentState);
		ExprContext _prevctx = _localctx;
		int _startState = 36;
		enterRecursionRule(_localctx, 36, RULE_expr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(448);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,39,_ctx) ) {
			case 1:
				{
				_localctx = new NegateContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(330);
				match(T__8);
				setState(331);
				expr(35);
				}
				break;
			case 2:
				{
				_localctx = new NotContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(332);
				match(T__9);
				setState(333);
				expr(34);
				}
				break;
			case 3:
				{
				_localctx = new AssignArithmeticContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(334);
				match(ID);
				setState(335);
				((AssignArithmeticContext)_localctx).op = _input.LT(1);
				_la = _input.LA(1);
				if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 31457280L) != 0)) ) {
					((AssignArithmeticContext)_localctx).op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(336);
				expr(27);
				}
				break;
			case 4:
				{
				_localctx = new NewInstanceContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(337);
				match(T__26);
				setState(338);
				match(ID);
				setState(339);
				match(L_PAREN);
				setState(341);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (((((_la - 9)) & ~0x3f) == 0 && ((1L << (_la - 9)) & 856055836861267971L) != 0)) {
					{
					setState(340);
					args();
					}
				}

				setState(343);
				match(R_PAREN);
				}
				break;
			case 5:
				{
				_localctx = new StructInstanceContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(344);
				match(ID);
				setState(345);
				match(L_CURLY);
				setState(346);
				instanceValues();
				setState(347);
				match(R_CURLY);
				}
				break;
			case 6:
				{
				_localctx = new SetValueMatrixContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(349);
				match(ID);
				setState(350);
				match(L_BRACKET);
				setState(351);
				expr(0);
				setState(352);
				match(R_BRACKET);
				setState(353);
				match(L_BRACKET);
				setState(354);
				expr(0);
				setState(355);
				match(R_BRACKET);
				setState(356);
				match(ASSIGN);
				setState(357);
				expr(23);
				}
				break;
			case 7:
				{
				_localctx = new GetValueMatrixContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(359);
				match(ID);
				setState(360);
				match(L_BRACKET);
				setState(361);
				expr(0);
				setState(362);
				match(R_BRACKET);
				setState(363);
				match(L_BRACKET);
				setState(364);
				expr(0);
				setState(366);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,37,_ctx) ) {
				case 1:
					{
					setState(365);
					match(R_BRACKET);
					}
					break;
				}
				}
				break;
			case 8:
				{
				_localctx = new SetValueIndexContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(368);
				match(ID);
				setState(369);
				match(L_BRACKET);
				setState(370);
				expr(0);
				setState(371);
				match(R_BRACKET);
				setState(372);
				match(ASSIGN);
				setState(373);
				((SetValueIndexContext)_localctx).valor = expr(21);
				}
				break;
			case 9:
				{
				_localctx = new GetValueIndexContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(375);
				match(ID);
				setState(376);
				match(L_BRACKET);
				setState(377);
				expr(0);
				setState(378);
				match(R_BRACKET);
				}
				break;
			case 10:
				{
				_localctx = new LenContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(380);
				match(LEN);
				setState(381);
				match(L_PAREN);
				setState(382);
				expr(0);
				setState(383);
				match(R_PAREN);
				}
				break;
			case 11:
				{
				_localctx = new StringJoinContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(385);
				match(JOIN);
				setState(386);
				match(L_PAREN);
				setState(387);
				match(ID);
				setState(388);
				match(T__0);
				setState(389);
				expr(0);
				setState(390);
				match(R_PAREN);
				}
				break;
			case 12:
				{
				_localctx = new AtoiContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(392);
				match(ATOI);
				setState(393);
				match(L_PAREN);
				setState(394);
				expr(0);
				setState(395);
				match(R_PAREN);
				}
				break;
			case 13:
				{
				_localctx = new ParseFloatContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(397);
				match(PARSEFLOAT);
				setState(398);
				match(L_PAREN);
				setState(399);
				expr(0);
				setState(400);
				match(R_PAREN);
				}
				break;
			case 14:
				{
				_localctx = new TypeOfContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(402);
				match(TYPEOF);
				setState(403);
				match(L_PAREN);
				setState(404);
				match(ID);
				setState(405);
				match(R_PAREN);
				}
				break;
			case 15:
				{
				_localctx = new SliceIndexContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(406);
				match(SLICE_INDEX);
				setState(407);
				match(L_PAREN);
				setState(408);
				match(ID);
				setState(409);
				match(T__0);
				setState(410);
				expr(0);
				setState(411);
				match(R_PAREN);
				}
				break;
			case 16:
				{
				_localctx = new IncrementDecrementContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(413);
				match(ID);
				setState(414);
				((IncrementDecrementContext)_localctx).op = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==T__27 || _la==T__28) ) {
					((IncrementDecrementContext)_localctx).op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
				break;
			case 17:
				{
				_localctx = new IdentifierContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(415);
				match(ID);
				}
				break;
			case 18:
				{
				_localctx = new RuneContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(416);
				match(RUNE);
				}
				break;
			case 19:
				{
				_localctx = new BooleanContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(417);
				match(BOOL);
				}
				break;
			case 20:
				{
				_localctx = new FloatContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(418);
				match(FLOAT);
				}
				break;
			case 21:
				{
				_localctx = new StringContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(419);
				match(STRING);
				}
				break;
			case 22:
				{
				_localctx = new IntContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(420);
				match(INT);
				}
				break;
			case 23:
				{
				_localctx = new SliceContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(421);
				match(L_CURLY);
				setState(425);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (((((_la - 9)) & ~0x3f) == 0 && ((1L << (_la - 9)) & 856055836861267971L) != 0)) {
					{
					{
					setState(422);
					sliceValues();
					}
					}
					setState(427);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(428);
				match(R_CURLY);
				}
				break;
			case 24:
				{
				_localctx = new SliceWithValuesContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(429);
				match(L_BRACKET);
				setState(430);
				match(R_BRACKET);
				setState(431);
				type();
				setState(432);
				match(L_CURLY);
				setState(433);
				sliceValues();
				setState(434);
				match(R_CURLY);
				}
				break;
			case 25:
				{
				_localctx = new AppendContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(436);
				match(APPEND);
				setState(437);
				match(L_PAREN);
				setState(438);
				match(ID);
				setState(439);
				match(T__0);
				setState(440);
				expr(0);
				setState(441);
				match(R_PAREN);
				}
				break;
			case 26:
				{
				_localctx = new NilContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(443);
				match(NIL);
				}
				break;
			case 27:
				{
				_localctx = new ParensContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(444);
				match(L_PAREN);
				setState(445);
				expr(0);
				setState(446);
				match(R_PAREN);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(479);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,42,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(477);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,41,_ctx) ) {
					case 1:
						{
						_localctx = new MulDivContext(new ExprContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(450);
						if (!(precpred(_ctx, 32))) throw new FailedPredicateException(this, "precpred(_ctx, 32)");
						setState(451);
						((MulDivContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__10 || _la==T__11) ) {
							((MulDivContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(452);
						expr(33);
						}
						break;
					case 2:
						{
						_localctx = new AddSubContext(new ExprContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(453);
						if (!(precpred(_ctx, 31))) throw new FailedPredicateException(this, "precpred(_ctx, 31)");
						setState(454);
						((AddSubContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__8 || _la==T__12) ) {
							((AddSubContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(455);
						expr(32);
						}
						break;
					case 3:
						{
						_localctx = new RelationalContext(new ExprContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(456);
						if (!(precpred(_ctx, 30))) throw new FailedPredicateException(this, "precpred(_ctx, 30)");
						setState(457);
						((RelationalContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 245760L) != 0)) ) {
							((RelationalContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(458);
						expr(31);
						}
						break;
					case 4:
						{
						_localctx = new ModContext(new ExprContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(459);
						if (!(precpred(_ctx, 29))) throw new FailedPredicateException(this, "precpred(_ctx, 29)");
						setState(460);
						match(T__17);
						setState(461);
						expr(30);
						}
						break;
					case 5:
						{
						_localctx = new EqualityContext(new ExprContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(462);
						if (!(precpred(_ctx, 28))) throw new FailedPredicateException(this, "precpred(_ctx, 28)");
						setState(463);
						((EqualityContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__18 || _la==T__19) ) {
							((EqualityContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(464);
						expr(29);
						}
						break;
					case 6:
						{
						_localctx = new LogicalContext(new ExprContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(465);
						if (!(precpred(_ctx, 26))) throw new FailedPredicateException(this, "precpred(_ctx, 26)");
						setState(466);
						((LogicalContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==T__24 || _la==T__25) ) {
							((LogicalContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(467);
						expr(27);
						}
						break;
					case 7:
						{
						_localctx = new AssignContext(new ExprContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(468);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(469);
						match(ASSIGN);
						setState(470);
						expr(14);
						}
						break;
					case 8:
						{
						_localctx = new CallStmtContext(new ExprContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(471);
						if (!(precpred(_ctx, 33))) throw new FailedPredicateException(this, "precpred(_ctx, 33)");
						setState(473); 
						_errHandler.sync(this);
						_alt = 1;
						do {
							switch (_alt) {
							case 1:
								{
								{
								setState(472);
								call();
								}
								}
								break;
							default:
								throw new NoViableAltException(this);
							}
							setState(475); 
							_errHandler.sync(this);
							_alt = getInterpreter().adaptivePredict(_input,40,_ctx);
						} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
						}
						break;
					}
					} 
				}
				setState(481);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,42,_ctx);
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
	public static class CallContext extends ParserRuleContext {
		public CallContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_call; }
	 
		public CallContext() { }
		public void copyFrom(CallContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FuncCallContext extends CallContext {
		public TerminalNode L_PAREN() { return getToken(V4LangGrammarParser.L_PAREN, 0); }
		public TerminalNode R_PAREN() { return getToken(V4LangGrammarParser.R_PAREN, 0); }
		public ArgsContext args() {
			return getRuleContext(ArgsContext.class,0);
		}
		public FuncCallContext(CallContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class GetContext extends CallContext {
		public TerminalNode ID() { return getToken(V4LangGrammarParser.ID, 0); }
		public GetContext(CallContext ctx) { copyFrom(ctx); }
	}

	public final CallContext call() throws RecognitionException {
		CallContext _localctx = new CallContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_call);
		int _la;
		try {
			setState(489);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case L_PAREN:
				_localctx = new FuncCallContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(482);
				match(L_PAREN);
				setState(484);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (((((_la - 9)) & ~0x3f) == 0 && ((1L << (_la - 9)) & 856055836861267971L) != 0)) {
					{
					setState(483);
					args();
					}
				}

				setState(486);
				match(R_PAREN);
				}
				break;
			case T__29:
				_localctx = new GetContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(487);
				match(T__29);
				setState(488);
				match(ID);
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
	public static class InstanceValuesContext extends ParserRuleContext {
		public List<TerminalNode> ID() { return getTokens(V4LangGrammarParser.ID); }
		public TerminalNode ID(int i) {
			return getToken(V4LangGrammarParser.ID, i);
		}
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public InstanceValuesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instanceValues; }
	}

	public final InstanceValuesContext instanceValues() throws RecognitionException {
		InstanceValuesContext _localctx = new InstanceValuesContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_instanceValues);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(491);
			match(ID);
			setState(492);
			match(T__6);
			setState(493);
			expr(0);
			setState(500);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,45,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(494);
					match(T__0);
					setState(495);
					match(ID);
					setState(496);
					match(T__6);
					setState(497);
					expr(0);
					}
					} 
				}
				setState(502);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,45,_ctx);
			}
			setState(504);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==T__0) {
				{
				setState(503);
				match(T__0);
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
	public static class TypeContext extends ParserRuleContext {
		public TypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_type; }
	}

	public final TypeContext type() throws RecognitionException {
		TypeContext _localctx = new TypeContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_type);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(506);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 66571993088L) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
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

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 18:
			return expr_sempred((ExprContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expr_sempred(ExprContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 32);
		case 1:
			return precpred(_ctx, 31);
		case 2:
			return precpred(_ctx, 30);
		case 3:
			return precpred(_ctx, 29);
		case 4:
			return precpred(_ctx, 28);
		case 5:
			return precpred(_ctx, 26);
		case 6:
			return precpred(_ctx, 13);
		case 7:
			return precpred(_ctx, 33);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001F\u01fd\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002\u0012\u0007\u0012"+
		"\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0002\u0015\u0007\u0015"+
		"\u0001\u0000\u0005\u0000.\b\u0000\n\u0000\f\u00001\t\u0000\u0001\u0001"+
		"\u0001\u0001\u0003\u00015\b\u0001\u0001\u0001\u0001\u0001\u0003\u0001"+
		"9\b\u0001\u0001\u0001\u0001\u0001\u0003\u0001=\b\u0001\u0001\u0001\u0001"+
		"\u0001\u0003\u0001A\b\u0001\u0001\u0001\u0001\u0001\u0003\u0001E\b\u0001"+
		"\u0001\u0001\u0001\u0001\u0003\u0001I\b\u0001\u0003\u0001K\b\u0001\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0003\u0002b\b\u0002\u0001\u0003\u0001"+
		"\u0003\u0001\u0003\u0001\u0003\u0003\u0003h\b\u0003\u0001\u0003\u0001"+
		"\u0003\u0003\u0003l\b\u0003\u0001\u0003\u0001\u0003\u0005\u0003p\b\u0003"+
		"\n\u0003\f\u0003s\t\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0001\u0003\u0001\u0003\u0003\u0003{\b\u0003\u0001\u0003\u0001\u0003"+
		"\u0003\u0003\u007f\b\u0003\u0001\u0003\u0001\u0003\u0005\u0003\u0083\b"+
		"\u0003\n\u0003\f\u0003\u0086\t\u0003\u0001\u0003\u0001\u0003\u0003\u0003"+
		"\u008a\b\u0003\u0001\u0004\u0001\u0004\u0001\u0004\u0003\u0004\u008f\b"+
		"\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0003"+
		"\u0004\u0096\b\u0004\u0001\u0004\u0005\u0004\u0099\b\u0004\n\u0004\f\u0004"+
		"\u009c\t\u0004\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0005\u0005\u00a3\b\u0005\n\u0005\f\u0005\u00a6\t\u0005\u0001\u0005\u0001"+
		"\u0005\u0001\u0006\u0001\u0006\u0003\u0006\u00ac\b\u0006\u0001\u0007\u0001"+
		"\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\b\u0001\b\u0001\b\u0001"+
		"\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0003"+
		"\b\u00bf\b\b\u0001\t\u0001\t\u0001\t\u0005\t\u00c4\b\t\n\t\f\t\u00c7\t"+
		"\t\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001\n\u0001"+
		"\n\u0001\n\u0001\n\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001"+
		"\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0005\u000b\u00dc\b\u000b\n"+
		"\u000b\f\u000b\u00df\t\u000b\u0001\u000b\u0003\u000b\u00e2\b\u000b\u0001"+
		"\f\u0001\f\u0001\f\u0001\f\u0003\f\u00e8\b\f\u0001\f\u0001\f\u0001\f\u0005"+
		"\f\u00ed\b\f\n\f\f\f\u00f0\t\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f"+
		"\u0005\f\u00f7\b\f\n\f\f\f\u00fa\t\f\u0001\f\u0003\f\u00fd\b\f\u0001\f"+
		"\u0001\f\u0001\f\u0001\f\u0005\f\u0103\b\f\n\f\f\f\u0106\t\f\u0001\f\u0003"+
		"\f\u0109\b\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001"+
		"\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001"+
		"\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001\f\u0001"+
		"\f\u0003\f\u0125\b\f\u0003\f\u0127\b\f\u0001\r\u0001\r\u0001\r\u0001\r"+
		"\u0001\r\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000f\u0001\u000f\u0001"+
		"\u000f\u0001\u000f\u0005\u000f\u0135\b\u000f\n\u000f\f\u000f\u0138\t\u000f"+
		"\u0001\u0010\u0001\u0010\u0001\u0010\u0005\u0010\u013d\b\u0010\n\u0010"+
		"\f\u0010\u0140\t\u0010\u0001\u0011\u0001\u0011\u0001\u0011\u0005\u0011"+
		"\u0145\b\u0011\n\u0011\f\u0011\u0148\t\u0011\u0001\u0012\u0001\u0012\u0001"+
		"\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001"+
		"\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0003\u0012\u0156\b\u0012\u0001"+
		"\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001"+
		"\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001"+
		"\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001"+
		"\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0003\u0012\u016f"+
		"\b\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001"+
		"\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001"+
		"\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001"+
		"\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001"+
		"\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001"+
		"\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001"+
		"\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001"+
		"\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001"+
		"\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001"+
		"\u0012\u0001\u0012\u0005\u0012\u01a8\b\u0012\n\u0012\f\u0012\u01ab\t\u0012"+
		"\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012"+
		"\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012"+
		"\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012"+
		"\u0001\u0012\u0001\u0012\u0003\u0012\u01c1\b\u0012\u0001\u0012\u0001\u0012"+
		"\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012"+
		"\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012"+
		"\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012\u0001\u0012"+
		"\u0001\u0012\u0001\u0012\u0001\u0012\u0004\u0012\u01da\b\u0012\u000b\u0012"+
		"\f\u0012\u01db\u0005\u0012\u01de\b\u0012\n\u0012\f\u0012\u01e1\t\u0012"+
		"\u0001\u0013\u0001\u0013\u0003\u0013\u01e5\b\u0013\u0001\u0013\u0001\u0013"+
		"\u0001\u0013\u0003\u0013\u01ea\b\u0013\u0001\u0014\u0001\u0014\u0001\u0014"+
		"\u0001\u0014\u0001\u0014\u0001\u0014\u0001\u0014\u0005\u0014\u01f3\b\u0014"+
		"\n\u0014\f\u0014\u01f6\t\u0014\u0001\u0014\u0003\u0014\u01f9\b\u0014\u0001"+
		"\u0015\u0001\u0015\u0001\u0015\u0000\u0001$\u0016\u0000\u0002\u0004\u0006"+
		"\b\n\f\u000e\u0010\u0012\u0014\u0016\u0018\u001a\u001c\u001e \"$&(*\u0000"+
		"\b\u0001\u0000\u0015\u0018\u0001\u0000\u001c\u001d\u0001\u0000\u000b\f"+
		"\u0002\u0000\t\t\r\r\u0001\u0000\u000e\u0011\u0001\u0000\u0013\u0014\u0001"+
		"\u0000\u0019\u001a\u0001\u0000\u001f#\u0245\u0000/\u0001\u0000\u0000\u0000"+
		"\u0002J\u0001\u0000\u0000\u0000\u0004a\u0001\u0000\u0000\u0000\u0006\u0089"+
		"\u0001\u0000\u0000\u0000\b\u008b\u0001\u0000\u0000\u0000\n\u009d\u0001"+
		"\u0000\u0000\u0000\f\u00ab\u0001\u0000\u0000\u0000\u000e\u00ad\u0001\u0000"+
		"\u0000\u0000\u0010\u00be\u0001\u0000\u0000\u0000\u0012\u00c0\u0001\u0000"+
		"\u0000\u0000\u0014\u00c8\u0001\u0000\u0000\u0000\u0016\u00d3\u0001\u0000"+
		"\u0000\u0000\u0018\u0126\u0001\u0000\u0000\u0000\u001a\u0128\u0001\u0000"+
		"\u0000\u0000\u001c\u012d\u0001\u0000\u0000\u0000\u001e\u0130\u0001\u0000"+
		"\u0000\u0000 \u0139\u0001\u0000\u0000\u0000\"\u0141\u0001\u0000\u0000"+
		"\u0000$\u01c0\u0001\u0000\u0000\u0000&\u01e9\u0001\u0000\u0000\u0000("+
		"\u01eb\u0001\u0000\u0000\u0000*\u01fa\u0001\u0000\u0000\u0000,.\u0003"+
		"\u0002\u0001\u0000-,\u0001\u0000\u0000\u0000.1\u0001\u0000\u0000\u0000"+
		"/-\u0001\u0000\u0000\u0000/0\u0001\u0000\u0000\u00000\u0001\u0001\u0000"+
		"\u0000\u00001/\u0001\u0000\u0000\u000024\u0003\u0004\u0002\u000035\u0005"+
		"=\u0000\u000043\u0001\u0000\u0000\u000045\u0001\u0000\u0000\u00005K\u0001"+
		"\u0000\u0000\u000068\u0003\u0014\n\u000079\u0005=\u0000\u000087\u0001"+
		"\u0000\u0000\u000089\u0001\u0000\u0000\u00009K\u0001\u0000\u0000\u0000"+
		":<\u0003\u0010\b\u0000;=\u0005=\u0000\u0000<;\u0001\u0000\u0000\u0000"+
		"<=\u0001\u0000\u0000\u0000=K\u0001\u0000\u0000\u0000>@\u0003\u0018\f\u0000"+
		"?A\u0005=\u0000\u0000@?\u0001\u0000\u0000\u0000@A\u0001\u0000\u0000\u0000"+
		"AK\u0001\u0000\u0000\u0000BD\u0003\u0006\u0003\u0000CE\u0005=\u0000\u0000"+
		"DC\u0001\u0000\u0000\u0000DE\u0001\u0000\u0000\u0000EK\u0001\u0000\u0000"+
		"\u0000FH\u0003\n\u0005\u0000GI\u0005=\u0000\u0000HG\u0001\u0000\u0000"+
		"\u0000HI\u0001\u0000\u0000\u0000IK\u0001\u0000\u0000\u0000J2\u0001\u0000"+
		"\u0000\u0000J6\u0001\u0000\u0000\u0000J:\u0001\u0000\u0000\u0000J>\u0001"+
		"\u0000\u0000\u0000JB\u0001\u0000\u0000\u0000JF\u0001\u0000\u0000\u0000"+
		"K\u0003\u0001\u0000\u0000\u0000LM\u0005$\u0000\u0000MN\u0005D\u0000\u0000"+
		"NO\u0003*\u0015\u0000OP\u0005<\u0000\u0000PQ\u0003$\u0012\u0000Qb\u0001"+
		"\u0000\u0000\u0000RS\u0005$\u0000\u0000ST\u0005D\u0000\u0000Tb\u0003*"+
		"\u0015\u0000UV\u0005$\u0000\u0000VW\u0005D\u0000\u0000WX\u00055\u0000"+
		"\u0000XY\u00056\u0000\u0000Yb\u0003*\u0015\u0000Z[\u0005D\u0000\u0000"+
		"[\\\u0005;\u0000\u0000\\b\u0003$\u0012\u0000]^\u0005D\u0000\u0000^b\u0003"+
		"*\u0015\u0000_`\u0005D\u0000\u0000`b\u0005D\u0000\u0000aL\u0001\u0000"+
		"\u0000\u0000aR\u0001\u0000\u0000\u0000aU\u0001\u0000\u0000\u0000aZ\u0001"+
		"\u0000\u0000\u0000a]\u0001\u0000\u0000\u0000a_\u0001\u0000\u0000\u0000"+
		"b\u0005\u0001\u0000\u0000\u0000cd\u00053\u0000\u0000de\u0005D\u0000\u0000"+
		"eg\u00057\u0000\u0000fh\u0003\b\u0004\u0000gf\u0001\u0000\u0000\u0000"+
		"gh\u0001\u0000\u0000\u0000hi\u0001\u0000\u0000\u0000ik\u00058\u0000\u0000"+
		"jl\u0003*\u0015\u0000kj\u0001\u0000\u0000\u0000kl\u0001\u0000\u0000\u0000"+
		"lm\u0001\u0000\u0000\u0000mq\u00055\u0000\u0000np\u0003\u0002\u0001\u0000"+
		"on\u0001\u0000\u0000\u0000ps\u0001\u0000\u0000\u0000qo\u0001\u0000\u0000"+
		"\u0000qr\u0001\u0000\u0000\u0000rt\u0001\u0000\u0000\u0000sq\u0001\u0000"+
		"\u0000\u0000t\u008a\u00056\u0000\u0000uv\u00053\u0000\u0000vw\u0003\u000e"+
		"\u0007\u0000wx\u0005D\u0000\u0000xz\u00057\u0000\u0000y{\u0003\b\u0004"+
		"\u0000zy\u0001\u0000\u0000\u0000z{\u0001\u0000\u0000\u0000{|\u0001\u0000"+
		"\u0000\u0000|~\u00058\u0000\u0000}\u007f\u0003*\u0015\u0000~}\u0001\u0000"+
		"\u0000\u0000~\u007f\u0001\u0000\u0000\u0000\u007f\u0080\u0001\u0000\u0000"+
		"\u0000\u0080\u0084\u00055\u0000\u0000\u0081\u0083\u0003\u0002\u0001\u0000"+
		"\u0082\u0081\u0001\u0000\u0000\u0000\u0083\u0086\u0001\u0000\u0000\u0000"+
		"\u0084\u0082\u0001\u0000\u0000\u0000\u0084\u0085\u0001\u0000\u0000\u0000"+
		"\u0085\u0087\u0001\u0000\u0000\u0000\u0086\u0084\u0001\u0000\u0000\u0000"+
		"\u0087\u0088\u00056\u0000\u0000\u0088\u008a\u0001\u0000\u0000\u0000\u0089"+
		"c\u0001\u0000\u0000\u0000\u0089u\u0001\u0000\u0000\u0000\u008a\u0007\u0001"+
		"\u0000\u0000\u0000\u008b\u008e\u0005D\u0000\u0000\u008c\u008d\u00059\u0000"+
		"\u0000\u008d\u008f\u0005:\u0000\u0000\u008e\u008c\u0001\u0000\u0000\u0000"+
		"\u008e\u008f\u0001\u0000\u0000\u0000\u008f\u0090\u0001\u0000\u0000\u0000"+
		"\u0090\u009a\u0003*\u0015\u0000\u0091\u0092\u0005\u0001\u0000\u0000\u0092"+
		"\u0095\u0005D\u0000\u0000\u0093\u0094\u00059\u0000\u0000\u0094\u0096\u0005"+
		":\u0000\u0000\u0095\u0093\u0001\u0000\u0000\u0000\u0095\u0096\u0001\u0000"+
		"\u0000\u0000\u0096\u0097\u0001\u0000\u0000\u0000\u0097\u0099\u0003*\u0015"+
		"\u0000\u0098\u0091\u0001\u0000\u0000\u0000\u0099\u009c\u0001\u0000\u0000"+
		"\u0000\u009a\u0098\u0001\u0000\u0000\u0000\u009a\u009b\u0001\u0000\u0000"+
		"\u0000\u009b\t\u0001\u0000\u0000\u0000\u009c\u009a\u0001\u0000\u0000\u0000"+
		"\u009d\u009e\u0005\u0002\u0000\u0000\u009e\u009f\u0005D\u0000\u0000\u009f"+
		"\u00a0\u00054\u0000\u0000\u00a0\u00a4\u00055\u0000\u0000\u00a1\u00a3\u0003"+
		"\f\u0006\u0000\u00a2\u00a1\u0001\u0000\u0000\u0000\u00a3\u00a6\u0001\u0000"+
		"\u0000\u0000\u00a4\u00a2\u0001\u0000\u0000\u0000\u00a4\u00a5\u0001\u0000"+
		"\u0000\u0000\u00a5\u00a7\u0001\u0000\u0000\u0000\u00a6\u00a4\u0001\u0000"+
		"\u0000\u0000\u00a7\u00a8\u00056\u0000\u0000\u00a8\u000b\u0001\u0000\u0000"+
		"\u0000\u00a9\u00ac\u0003\u0004\u0002\u0000\u00aa\u00ac\u0003\u0006\u0003"+
		"\u0000\u00ab\u00a9\u0001\u0000\u0000\u0000\u00ab\u00aa\u0001\u0000\u0000"+
		"\u0000\u00ac\r\u0001\u0000\u0000\u0000\u00ad\u00ae\u00057\u0000\u0000"+
		"\u00ae\u00af\u0005D\u0000\u0000\u00af\u00b0\u0005D\u0000\u0000\u00b0\u00b1"+
		"\u00058\u0000\u0000\u00b1\u000f\u0001\u0000\u0000\u0000\u00b2\u00b3\u0005"+
		"$\u0000\u0000\u00b3\u00b4\u0005D\u0000\u0000\u00b4\u00b5\u0005;\u0000"+
		"\u0000\u00b5\u00b6\u00059\u0000\u0000\u00b6\u00b7\u0005:\u0000\u0000\u00b7"+
		"\u00b8\u0003*\u0015\u0000\u00b8\u00b9\u00055\u0000\u0000\u00b9\u00ba\u0003"+
		"\u0012\t\u0000\u00ba\u00bb\u00056\u0000\u0000\u00bb\u00bf\u0001\u0000"+
		"\u0000\u0000\u00bc\u00bd\u0005D\u0000\u0000\u00bd\u00bf\u0003*\u0015\u0000"+
		"\u00be\u00b2\u0001\u0000\u0000\u0000\u00be\u00bc\u0001\u0000\u0000\u0000"+
		"\u00bf\u0011\u0001\u0000\u0000\u0000\u00c0\u00c5\u0003$\u0012\u0000\u00c1"+
		"\u00c2\u0005\u0001\u0000\u0000\u00c2\u00c4\u0003$\u0012\u0000\u00c3\u00c1"+
		"\u0001\u0000\u0000\u0000\u00c4\u00c7\u0001\u0000\u0000\u0000\u00c5\u00c3"+
		"\u0001\u0000\u0000\u0000\u00c5\u00c6\u0001\u0000\u0000\u0000\u00c6\u0013"+
		"\u0001\u0000\u0000\u0000\u00c7\u00c5\u0001\u0000\u0000\u0000\u00c8\u00c9"+
		"\u0005D\u0000\u0000\u00c9\u00ca\u0005;\u0000\u0000\u00ca\u00cb\u00059"+
		"\u0000\u0000\u00cb\u00cc\u0005:\u0000\u0000\u00cc\u00cd\u00059\u0000\u0000"+
		"\u00cd\u00ce\u0005:\u0000\u0000\u00ce\u00cf\u0003*\u0015\u0000\u00cf\u00d0"+
		"\u00055\u0000\u0000\u00d0\u00d1\u0003\u0016\u000b\u0000\u00d1\u00d2\u0005"+
		"6\u0000\u0000\u00d2\u0015\u0001\u0000\u0000\u0000\u00d3\u00d4\u00055\u0000"+
		"\u0000\u00d4\u00d5\u0003\u0012\t\u0000\u00d5\u00dd\u00056\u0000\u0000"+
		"\u00d6\u00d7\u0005\u0001\u0000\u0000\u00d7\u00d8\u00055\u0000\u0000\u00d8"+
		"\u00d9\u0003\u0012\t\u0000\u00d9\u00da\u00056\u0000\u0000\u00da\u00dc"+
		"\u0001\u0000\u0000\u0000\u00db\u00d6\u0001\u0000\u0000\u0000\u00dc\u00df"+
		"\u0001\u0000\u0000\u0000\u00dd\u00db\u0001\u0000\u0000\u0000\u00dd\u00de"+
		"\u0001\u0000\u0000\u0000\u00de\u00e1\u0001\u0000\u0000\u0000\u00df\u00dd"+
		"\u0001\u0000\u0000\u0000\u00e0\u00e2\u0005\u0001\u0000\u0000\u00e1\u00e0"+
		"\u0001\u0000\u0000\u0000\u00e1\u00e2\u0001\u0000\u0000\u0000\u00e2\u0017"+
		"\u0001\u0000\u0000\u0000\u00e3\u0127\u0003$\u0012\u0000\u00e4\u00e5\u0005"+
		"%\u0000\u0000\u00e5\u00e7\u00057\u0000\u0000\u00e6\u00e8\u0003\"\u0011"+
		"\u0000\u00e7\u00e6\u0001\u0000\u0000\u0000\u00e7\u00e8\u0001\u0000\u0000"+
		"\u0000\u00e8\u00e9\u0001\u0000\u0000\u0000\u00e9\u0127\u00058\u0000\u0000"+
		"\u00ea\u00ee\u00055\u0000\u0000\u00eb\u00ed\u0003\u0002\u0001\u0000\u00ec"+
		"\u00eb\u0001\u0000\u0000\u0000\u00ed\u00f0\u0001\u0000\u0000\u0000\u00ee"+
		"\u00ec\u0001\u0000\u0000\u0000\u00ee\u00ef\u0001\u0000\u0000\u0000\u00ef"+
		"\u00f1\u0001\u0000\u0000\u0000\u00f0\u00ee\u0001\u0000\u0000\u0000\u00f1"+
		"\u0127\u00056\u0000\u0000\u00f2\u00f3\u0005&\u0000\u0000\u00f3\u00f4\u0003"+
		"$\u0012\u0000\u00f4\u00f8\u0003\u0018\f\u0000\u00f5\u00f7\u0003\u001a"+
		"\r\u0000\u00f6\u00f5\u0001\u0000\u0000\u0000\u00f7\u00fa\u0001\u0000\u0000"+
		"\u0000\u00f8\u00f6\u0001\u0000\u0000\u0000\u00f8\u00f9\u0001\u0000\u0000"+
		"\u0000\u00f9\u00fc\u0001\u0000\u0000\u0000\u00fa\u00f8\u0001\u0000\u0000"+
		"\u0000\u00fb\u00fd\u0003\u001c\u000e\u0000\u00fc\u00fb\u0001\u0000\u0000"+
		"\u0000\u00fc\u00fd\u0001\u0000\u0000\u0000\u00fd\u0127\u0001\u0000\u0000"+
		"\u0000\u00fe\u00ff\u0005\u0003\u0000\u0000\u00ff\u0100\u0003$\u0012\u0000"+
		"\u0100\u0104\u00055\u0000\u0000\u0101\u0103\u0003\u001e\u000f\u0000\u0102"+
		"\u0101\u0001\u0000\u0000\u0000\u0103\u0106\u0001\u0000\u0000\u0000\u0104"+
		"\u0102\u0001\u0000\u0000\u0000\u0104\u0105\u0001\u0000\u0000\u0000\u0105"+
		"\u0108\u0001\u0000\u0000\u0000\u0106\u0104\u0001\u0000\u0000\u0000\u0107"+
		"\u0109\u0003 \u0010\u0000\u0108\u0107\u0001\u0000\u0000\u0000\u0108\u0109"+
		"\u0001\u0000\u0000\u0000\u0109\u010a\u0001\u0000\u0000\u0000\u010a\u010b"+
		"\u00056\u0000\u0000\u010b\u0127\u0001\u0000\u0000\u0000\u010c\u010d\u0005"+
		"/\u0000\u0000\u010d\u010e\u0003\u0004\u0002\u0000\u010e\u010f\u0005=\u0000"+
		"\u0000\u010f\u0110\u0003$\u0012\u0000\u0110\u0111\u0005=\u0000\u0000\u0111"+
		"\u0112\u0003$\u0012\u0000\u0112\u0113\u0003\u0018\f\u0000\u0113\u0127"+
		"\u0001\u0000\u0000\u0000\u0114\u0115\u0005/\u0000\u0000\u0115\u0116\u0003"+
		"$\u0012\u0000\u0116\u0117\u0003\u0018\f\u0000\u0117\u0127\u0001\u0000"+
		"\u0000\u0000\u0118\u0119\u0005/\u0000\u0000\u0119\u011a\u0005D\u0000\u0000"+
		"\u011a\u011b\u0005\u0001\u0000\u0000\u011b\u011c\u0005D\u0000\u0000\u011c"+
		"\u011d\u0005;\u0000\u0000\u011d\u011e\u00050\u0000\u0000\u011e\u011f\u0005"+
		"D\u0000\u0000\u011f\u0127\u0003\u0018\f\u0000\u0120\u0127\u00051\u0000"+
		"\u0000\u0121\u0127\u0005\u0004\u0000\u0000\u0122\u0124\u0005\u0005\u0000"+
		"\u0000\u0123\u0125\u0003$\u0012\u0000\u0124\u0123\u0001\u0000\u0000\u0000"+
		"\u0124\u0125\u0001\u0000\u0000\u0000\u0125\u0127\u0001\u0000\u0000\u0000"+
		"\u0126\u00e3\u0001\u0000\u0000\u0000\u0126\u00e4\u0001\u0000\u0000\u0000"+
		"\u0126\u00ea\u0001\u0000\u0000\u0000\u0126\u00f2\u0001\u0000\u0000\u0000"+
		"\u0126\u00fe\u0001\u0000\u0000\u0000\u0126\u010c\u0001\u0000\u0000\u0000"+
		"\u0126\u0114\u0001\u0000\u0000\u0000\u0126\u0118\u0001\u0000\u0000\u0000"+
		"\u0126\u0120\u0001\u0000\u0000\u0000\u0126\u0121\u0001\u0000\u0000\u0000"+
		"\u0126\u0122\u0001\u0000\u0000\u0000\u0127\u0019\u0001\u0000\u0000\u0000"+
		"\u0128\u0129\u0005\'\u0000\u0000\u0129\u012a\u0005&\u0000\u0000\u012a"+
		"\u012b\u0003$\u0012\u0000\u012b\u012c\u0003\u0018\f\u0000\u012c\u001b"+
		"\u0001\u0000\u0000\u0000\u012d\u012e\u0005\'\u0000\u0000\u012e\u012f\u0003"+
		"\u0018\f\u0000\u012f\u001d\u0001\u0000\u0000\u0000\u0130\u0131\u0005\u0006"+
		"\u0000\u0000\u0131\u0132\u0003$\u0012\u0000\u0132\u0136\u0005\u0007\u0000"+
		"\u0000\u0133\u0135\u0003\u0002\u0001\u0000\u0134\u0133\u0001\u0000\u0000"+
		"\u0000\u0135\u0138\u0001\u0000\u0000\u0000\u0136\u0134\u0001\u0000\u0000"+
		"\u0000\u0136\u0137\u0001\u0000\u0000\u0000\u0137\u001f\u0001\u0000\u0000"+
		"\u0000\u0138\u0136\u0001\u0000\u0000\u0000\u0139\u013a\u0005\b\u0000\u0000"+
		"\u013a\u013e\u0005\u0007\u0000\u0000\u013b\u013d\u0003\u0002\u0001\u0000"+
		"\u013c\u013b\u0001\u0000\u0000\u0000\u013d\u0140\u0001\u0000\u0000\u0000"+
		"\u013e\u013c\u0001\u0000\u0000\u0000\u013e\u013f\u0001\u0000\u0000\u0000"+
		"\u013f!\u0001\u0000\u0000\u0000\u0140\u013e\u0001\u0000\u0000\u0000\u0141"+
		"\u0146\u0003$\u0012\u0000\u0142\u0143\u0005\u0001\u0000\u0000\u0143\u0145"+
		"\u0003$\u0012\u0000\u0144\u0142\u0001\u0000\u0000\u0000\u0145\u0148\u0001"+
		"\u0000\u0000\u0000\u0146\u0144\u0001\u0000\u0000\u0000\u0146\u0147\u0001"+
		"\u0000\u0000\u0000\u0147#\u0001\u0000\u0000\u0000\u0148\u0146\u0001\u0000"+
		"\u0000\u0000\u0149\u014a\u0006\u0012\uffff\uffff\u0000\u014a\u014b\u0005"+
		"\t\u0000\u0000\u014b\u01c1\u0003$\u0012#\u014c\u014d\u0005\n\u0000\u0000"+
		"\u014d\u01c1\u0003$\u0012\"\u014e\u014f\u0005D\u0000\u0000\u014f\u0150"+
		"\u0007\u0000\u0000\u0000\u0150\u01c1\u0003$\u0012\u001b\u0151\u0152\u0005"+
		"\u001b\u0000\u0000\u0152\u0153\u0005D\u0000\u0000\u0153\u0155\u00057\u0000"+
		"\u0000\u0154\u0156\u0003\"\u0011\u0000\u0155\u0154\u0001\u0000\u0000\u0000"+
		"\u0155\u0156\u0001\u0000\u0000\u0000\u0156\u0157\u0001\u0000\u0000\u0000"+
		"\u0157\u01c1\u00058\u0000\u0000\u0158\u0159\u0005D\u0000\u0000\u0159\u015a"+
		"\u00055\u0000\u0000\u015a\u015b\u0003(\u0014\u0000\u015b\u015c\u00056"+
		"\u0000\u0000\u015c\u01c1\u0001\u0000\u0000\u0000\u015d\u015e\u0005D\u0000"+
		"\u0000\u015e\u015f\u00059\u0000\u0000\u015f\u0160\u0003$\u0012\u0000\u0160"+
		"\u0161\u0005:\u0000\u0000\u0161\u0162\u00059\u0000\u0000\u0162\u0163\u0003"+
		"$\u0012\u0000\u0163\u0164\u0005:\u0000\u0000\u0164\u0165\u0005<\u0000"+
		"\u0000\u0165\u0166\u0003$\u0012\u0017\u0166\u01c1\u0001\u0000\u0000\u0000"+
		"\u0167\u0168\u0005D\u0000\u0000\u0168\u0169\u00059\u0000\u0000\u0169\u016a"+
		"\u0003$\u0012\u0000\u016a\u016b\u0005:\u0000\u0000\u016b\u016c\u00059"+
		"\u0000\u0000\u016c\u016e\u0003$\u0012\u0000\u016d\u016f\u0005:\u0000\u0000"+
		"\u016e\u016d\u0001\u0000\u0000\u0000\u016e\u016f\u0001\u0000\u0000\u0000"+
		"\u016f\u01c1\u0001\u0000\u0000\u0000\u0170\u0171\u0005D\u0000\u0000\u0171"+
		"\u0172\u00059\u0000\u0000\u0172\u0173\u0003$\u0012\u0000\u0173\u0174\u0005"+
		":\u0000\u0000\u0174\u0175\u0005<\u0000\u0000\u0175\u0176\u0003$\u0012"+
		"\u0015\u0176\u01c1\u0001\u0000\u0000\u0000\u0177\u0178\u0005D\u0000\u0000"+
		"\u0178\u0179\u00059\u0000\u0000\u0179\u017a\u0003$\u0012\u0000\u017a\u017b"+
		"\u0005:\u0000\u0000\u017b\u01c1\u0001\u0000\u0000\u0000\u017c\u017d\u0005"+
		")\u0000\u0000\u017d\u017e\u00057\u0000\u0000\u017e\u017f\u0003$\u0012"+
		"\u0000\u017f\u0180\u00058\u0000\u0000\u0180\u01c1\u0001\u0000\u0000\u0000"+
		"\u0181\u0182\u0005+\u0000\u0000\u0182\u0183\u00057\u0000\u0000\u0183\u0184"+
		"\u0005D\u0000\u0000\u0184\u0185\u0005\u0001\u0000\u0000\u0185\u0186\u0003"+
		"$\u0012\u0000\u0186\u0187\u00058\u0000\u0000\u0187\u01c1\u0001\u0000\u0000"+
		"\u0000\u0188\u0189\u0005,\u0000\u0000\u0189\u018a\u00057\u0000\u0000\u018a"+
		"\u018b\u0003$\u0012\u0000\u018b\u018c\u00058\u0000\u0000\u018c\u01c1\u0001"+
		"\u0000\u0000\u0000\u018d\u018e\u0005-\u0000\u0000\u018e\u018f\u00057\u0000"+
		"\u0000\u018f\u0190\u0003$\u0012\u0000\u0190\u0191\u00058\u0000\u0000\u0191"+
		"\u01c1\u0001\u0000\u0000\u0000\u0192\u0193\u0005.\u0000\u0000\u0193\u0194"+
		"\u00057\u0000\u0000\u0194\u0195\u0005D\u0000\u0000\u0195\u01c1\u00058"+
		"\u0000\u0000\u0196\u0197\u0005*\u0000\u0000\u0197\u0198\u00057\u0000\u0000"+
		"\u0198\u0199\u0005D\u0000\u0000\u0199\u019a\u0005\u0001\u0000\u0000\u019a"+
		"\u019b\u0003$\u0012\u0000\u019b\u019c\u00058\u0000\u0000\u019c\u01c1\u0001"+
		"\u0000\u0000\u0000\u019d\u019e\u0005D\u0000\u0000\u019e\u01c1\u0007\u0001"+
		"\u0000\u0000\u019f\u01c1\u0005D\u0000\u0000\u01a0\u01c1\u0005B\u0000\u0000"+
		"\u01a1\u01c1\u0005?\u0000\u0000\u01a2\u01c1\u0005@\u0000\u0000\u01a3\u01c1"+
		"\u0005A\u0000\u0000\u01a4\u01c1\u0005>\u0000\u0000\u01a5\u01a9\u00055"+
		"\u0000\u0000\u01a6\u01a8\u0003\u0012\t\u0000\u01a7\u01a6\u0001\u0000\u0000"+
		"\u0000\u01a8\u01ab\u0001\u0000\u0000\u0000\u01a9\u01a7\u0001\u0000\u0000"+
		"\u0000\u01a9\u01aa\u0001\u0000\u0000\u0000\u01aa\u01ac\u0001\u0000\u0000"+
		"\u0000\u01ab\u01a9\u0001\u0000\u0000\u0000\u01ac\u01c1\u00056\u0000\u0000"+
		"\u01ad\u01ae\u00059\u0000\u0000\u01ae\u01af\u0005:\u0000\u0000\u01af\u01b0"+
		"\u0003*\u0015\u0000\u01b0\u01b1\u00055\u0000\u0000\u01b1\u01b2\u0003\u0012"+
		"\t\u0000\u01b2\u01b3\u00056\u0000\u0000\u01b3\u01c1\u0001\u0000\u0000"+
		"\u0000\u01b4\u01b5\u0005(\u0000\u0000\u01b5\u01b6\u00057\u0000\u0000\u01b6"+
		"\u01b7\u0005D\u0000\u0000\u01b7\u01b8\u0005\u0001\u0000\u0000\u01b8\u01b9"+
		"\u0003$\u0012\u0000\u01b9\u01ba\u00058\u0000\u0000\u01ba\u01c1\u0001\u0000"+
		"\u0000\u0000\u01bb\u01c1\u00052\u0000\u0000\u01bc\u01bd\u00057\u0000\u0000"+
		"\u01bd\u01be\u0003$\u0012\u0000\u01be\u01bf\u00058\u0000\u0000\u01bf\u01c1"+
		"\u0001\u0000\u0000\u0000\u01c0\u0149\u0001\u0000\u0000\u0000\u01c0\u014c"+
		"\u0001\u0000\u0000\u0000\u01c0\u014e\u0001\u0000\u0000\u0000\u01c0\u0151"+
		"\u0001\u0000\u0000\u0000\u01c0\u0158\u0001\u0000\u0000\u0000\u01c0\u015d"+
		"\u0001\u0000\u0000\u0000\u01c0\u0167\u0001\u0000\u0000\u0000\u01c0\u0170"+
		"\u0001\u0000\u0000\u0000\u01c0\u0177\u0001\u0000\u0000\u0000\u01c0\u017c"+
		"\u0001\u0000\u0000\u0000\u01c0\u0181\u0001\u0000\u0000\u0000\u01c0\u0188"+
		"\u0001\u0000\u0000\u0000\u01c0\u018d\u0001\u0000\u0000\u0000\u01c0\u0192"+
		"\u0001\u0000\u0000\u0000\u01c0\u0196\u0001\u0000\u0000\u0000\u01c0\u019d"+
		"\u0001\u0000\u0000\u0000\u01c0\u019f\u0001\u0000\u0000\u0000\u01c0\u01a0"+
		"\u0001\u0000\u0000\u0000\u01c0\u01a1\u0001\u0000\u0000\u0000\u01c0\u01a2"+
		"\u0001\u0000\u0000\u0000\u01c0\u01a3\u0001\u0000\u0000\u0000\u01c0\u01a4"+
		"\u0001\u0000\u0000\u0000\u01c0\u01a5\u0001\u0000\u0000\u0000\u01c0\u01ad"+
		"\u0001\u0000\u0000\u0000\u01c0\u01b4\u0001\u0000\u0000\u0000\u01c0\u01bb"+
		"\u0001\u0000\u0000\u0000\u01c0\u01bc\u0001\u0000\u0000\u0000\u01c1\u01df"+
		"\u0001\u0000\u0000\u0000\u01c2\u01c3\n \u0000\u0000\u01c3\u01c4\u0007"+
		"\u0002\u0000\u0000\u01c4\u01de\u0003$\u0012!\u01c5\u01c6\n\u001f\u0000"+
		"\u0000\u01c6\u01c7\u0007\u0003\u0000\u0000\u01c7\u01de\u0003$\u0012 \u01c8"+
		"\u01c9\n\u001e\u0000\u0000\u01c9\u01ca\u0007\u0004\u0000\u0000\u01ca\u01de"+
		"\u0003$\u0012\u001f\u01cb\u01cc\n\u001d\u0000\u0000\u01cc\u01cd\u0005"+
		"\u0012\u0000\u0000\u01cd\u01de\u0003$\u0012\u001e\u01ce\u01cf\n\u001c"+
		"\u0000\u0000\u01cf\u01d0\u0007\u0005\u0000\u0000\u01d0\u01de\u0003$\u0012"+
		"\u001d\u01d1\u01d2\n\u001a\u0000\u0000\u01d2\u01d3\u0007\u0006\u0000\u0000"+
		"\u01d3\u01de\u0003$\u0012\u001b\u01d4\u01d5\n\r\u0000\u0000\u01d5\u01d6"+
		"\u0005<\u0000\u0000\u01d6\u01de\u0003$\u0012\u000e\u01d7\u01d9\n!\u0000"+
		"\u0000\u01d8\u01da\u0003&\u0013\u0000\u01d9\u01d8\u0001\u0000\u0000\u0000"+
		"\u01da\u01db\u0001\u0000\u0000\u0000\u01db\u01d9\u0001\u0000\u0000\u0000"+
		"\u01db\u01dc\u0001\u0000\u0000\u0000\u01dc\u01de\u0001\u0000\u0000\u0000"+
		"\u01dd\u01c2\u0001\u0000\u0000\u0000\u01dd\u01c5\u0001\u0000\u0000\u0000"+
		"\u01dd\u01c8\u0001\u0000\u0000\u0000\u01dd\u01cb\u0001\u0000\u0000\u0000"+
		"\u01dd\u01ce\u0001\u0000\u0000\u0000\u01dd\u01d1\u0001\u0000\u0000\u0000"+
		"\u01dd\u01d4\u0001\u0000\u0000\u0000\u01dd\u01d7\u0001\u0000\u0000\u0000"+
		"\u01de\u01e1\u0001\u0000\u0000\u0000\u01df\u01dd\u0001\u0000\u0000\u0000"+
		"\u01df\u01e0\u0001\u0000\u0000\u0000\u01e0%\u0001\u0000\u0000\u0000\u01e1"+
		"\u01df\u0001\u0000\u0000\u0000\u01e2\u01e4\u00057\u0000\u0000\u01e3\u01e5"+
		"\u0003\"\u0011\u0000\u01e4\u01e3\u0001\u0000\u0000\u0000\u01e4\u01e5\u0001"+
		"\u0000\u0000\u0000\u01e5\u01e6\u0001\u0000\u0000\u0000\u01e6\u01ea\u0005"+
		"8\u0000\u0000\u01e7\u01e8\u0005\u001e\u0000\u0000\u01e8\u01ea\u0005D\u0000"+
		"\u0000\u01e9\u01e2\u0001\u0000\u0000\u0000\u01e9\u01e7\u0001\u0000\u0000"+
		"\u0000\u01ea\'\u0001\u0000\u0000\u0000\u01eb\u01ec\u0005D\u0000\u0000"+
		"\u01ec\u01ed\u0005\u0007\u0000\u0000\u01ed\u01f4\u0003$\u0012\u0000\u01ee"+
		"\u01ef\u0005\u0001\u0000\u0000\u01ef\u01f0\u0005D\u0000\u0000\u01f0\u01f1"+
		"\u0005\u0007\u0000\u0000\u01f1\u01f3\u0003$\u0012\u0000\u01f2\u01ee\u0001"+
		"\u0000\u0000\u0000\u01f3\u01f6\u0001\u0000\u0000\u0000\u01f4\u01f2\u0001"+
		"\u0000\u0000\u0000\u01f4\u01f5\u0001\u0000\u0000\u0000\u01f5\u01f8\u0001"+
		"\u0000\u0000\u0000\u01f6\u01f4\u0001\u0000\u0000\u0000\u01f7\u01f9\u0005"+
		"\u0001\u0000\u0000\u01f8\u01f7\u0001\u0000\u0000\u0000\u01f8\u01f9\u0001"+
		"\u0000\u0000\u0000\u01f9)\u0001\u0000\u0000\u0000\u01fa\u01fb\u0007\u0007"+
		"\u0000\u0000\u01fb+\u0001\u0000\u0000\u0000//48<@DHJagkqz~\u0084\u0089"+
		"\u008e\u0095\u009a\u00a4\u00ab\u00be\u00c5\u00dd\u00e1\u00e7\u00ee\u00f8"+
		"\u00fc\u0104\u0108\u0124\u0126\u0136\u013e\u0146\u0155\u016e\u01a9\u01c0"+
		"\u01db\u01dd\u01df\u01e4\u01e9\u01f4\u01f8";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}