#include "options.h"        /* Compilation switches                   */
#include "cnst_fx.h"        /* Decoder static structure               */
#include "stl.h"

/* math_op.c */
const Word32 L_table_isqrt[48] =
{
     2147418112L,  2083389440L,  2024669184L,  1970667520L,
     1920794624L,  1874460672L,  1831403520L,  1791098880L,
     1753415680L,  1717960704L,  1684602880L,  1653145600L,
     1623326720L,  1595080704L,  1568276480L,  1542782976L,
     1518469120L,  1495334912L,  1473183744L,  1451950080L,
     1431633920L,  1412169728L,  1393491968L,  1375469568L,
     1358168064L,  1341521920L,  1325465600L,  1309933568L,
     1294991360L,  1280507904L,  1266548736L,  1252982784L,
     1239875584L,  1227161600L,  1214775296L,  1202847744L,
     1191182336L,  1179910144L,  1168965632L,  1158283264L,
     1147863040L,  1137770496L,  1127940096L,  1118306304L,
     1108934656L,  1099825152L,  1090912256L,  1082261504L
};

/* table of table_isqrt[i] - table_isqrt[i+1] */
const Word16 table_isqrt_diff[48] =
{
      977,   896,   824,   761,   707,   657,   615,   575,
      541,   509,   480,   455,   431,   409,   389,   371,
      353,   338,   324,   310,   297,   285,   275,   264,
      254,   245,   237,   228,   221,   213,   207,   200,
      194,   189,   182,   178,   172,   167,   163,   159,
      154,   150,   147,   143,   139,   136,   132,   130
};

const Word16 shift_Isqrt_lc[] = {9,10};

const Word16 table_pow2[32] =
{
    16384, 16743, 17109, 17484, 17867, 18258, 18658, 19066, 19484, 19911,
    20347, 20792, 21247, 21713, 22188, 22674, 23170, 23678, 24196, 24726,
    25268, 25821, 26386, 26964, 27554, 28158, 28774, 29405, 30048, 30706,
    31379, 32066
};

/* table of table_pow2[i+1] - table_pow2[i] */
const Word16 table_pow2_diff_x32[32] =
{
    11488, 11712, 12000, 12256, 12512, 12800, 13056, 13376, 13664, 13952,
    14240, 14560, 14912, 15200, 15552, 15872, 16256, 16576, 16960, 17344,
    17696, 18080, 18496, 18880, 19328, 19712, 20192, 20576, 21056, 21536,
    21984, 22432
};

const Word16 sqrt_table[49] =
{
    16384, 16888, 17378, 17854, 18318, 18770, 19212,
    19644, 20066, 20480, 20886, 21283, 21674, 22058,
    22435, 22806, 23170, 23530, 23884, 24232, 24576,
    24915, 25249, 25580, 25905, 26227, 26545, 26859,
    27170, 27477, 27780, 28081, 28378, 28672, 28963,
    29251, 29537, 29819, 30099, 30377, 30652, 30924,
    31194, 31462, 31727, 31991, 32252, 32511, 32767
};

/* log2.c */
const Word32 L_table_Log2_norm_lc[32] =
{
        -32768L,   95322112L,  187793408L,  277577728L,
     364871680L,  449740800L,  532381696L,  612859904L,
     691306496L,  767787008L,  842432512L,  915308544L,
     986546176L, 1056210944L, 1124302848L, 1190887424L,
    1256095744L, 1319993344L, 1382580224L, 1443921920L,
    1504083968L, 1563131904L, 1621000192L, 1677885440L,
    1733722112L, 1788510208L, 1842380800L, 1895399424L,
    1947435008L, 1998618624L, 2049015808L, 2098626560L
};
 
const Word16 table_diff_Log2_norm_lc[32] =
{
     1455, 1411,  1370, 1332,  1295, 1261,  1228, 1197,
     1167, 1139,  1112, 1087,  1063, 1039,  1016,  995,
      975,  955,   936,  918,   901,  883,   868,  852,
      836,  822,   809,  794,   781,  769,   757,  744
};
 
const Word16 log2_tab[33]={
    0x7800, 0x782D, 0x785A, 0x7884, 0x78AE, 0x78D6, 0x78FE, 0x7924,
    0x794A, 0x796E, 0x7992, 0x79B4, 0x79D6, 0x79F8, 0x7A18, 0x7A38,
    0x7A57, 0x7A75, 0x7A93, 0x7AB1, 0x7ACD, 0x7AE9, 0x7B05, 0x7B20,
    0x7B3B, 0x7B55, 0x7B6F, 0x7B88, 0x7BA1, 0x7BB9, 0x7BD1, 0x7BE9, 
    0x7C00
};
