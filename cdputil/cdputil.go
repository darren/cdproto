package cdputil

// Code generated by chromedp-gen. DO NOT EDIT.

import (
	"errors"

	cdp "github.com/knq/chromedp/cdp"
	"github.com/knq/chromedp/cdp/accessibility"
	"github.com/knq/chromedp/cdp/animation"
	"github.com/knq/chromedp/cdp/applicationcache"
	"github.com/knq/chromedp/cdp/audits"
	"github.com/knq/chromedp/cdp/browser"
	"github.com/knq/chromedp/cdp/cachestorage"
	"github.com/knq/chromedp/cdp/css"
	"github.com/knq/chromedp/cdp/database"
	"github.com/knq/chromedp/cdp/debugger"
	"github.com/knq/chromedp/cdp/dom"
	"github.com/knq/chromedp/cdp/domdebugger"
	"github.com/knq/chromedp/cdp/domsnapshot"
	"github.com/knq/chromedp/cdp/domstorage"
	"github.com/knq/chromedp/cdp/emulation"
	"github.com/knq/chromedp/cdp/headlessexperimental"
	"github.com/knq/chromedp/cdp/heapprofiler"
	"github.com/knq/chromedp/cdp/indexeddb"
	"github.com/knq/chromedp/cdp/inspector"
	iodom "github.com/knq/chromedp/cdp/io"
	"github.com/knq/chromedp/cdp/layertree"
	logdom "github.com/knq/chromedp/cdp/log"
	"github.com/knq/chromedp/cdp/memory"
	"github.com/knq/chromedp/cdp/network"
	"github.com/knq/chromedp/cdp/overlay"
	"github.com/knq/chromedp/cdp/page"
	"github.com/knq/chromedp/cdp/performance"
	"github.com/knq/chromedp/cdp/profiler"
	"github.com/knq/chromedp/cdp/runtime"
	"github.com/knq/chromedp/cdp/schema"
	"github.com/knq/chromedp/cdp/security"
	"github.com/knq/chromedp/cdp/serviceworker"
	"github.com/knq/chromedp/cdp/storage"
	"github.com/knq/chromedp/cdp/systeminfo"
	"github.com/knq/chromedp/cdp/target"
	"github.com/knq/chromedp/cdp/tethering"
	"github.com/knq/chromedp/cdp/tracing"
	"github.com/mailru/easyjson"
)

type empty struct{}

var emptyVal = &empty{}

// UnmarshalMessage unmarshals the message result or params.
func UnmarshalMessage(msg *cdp.Message) (interface{}, error) {
	var v easyjson.Unmarshaler
	switch msg.Method {
	case cdp.CommandInspectorEnable:
		return emptyVal, nil

	case cdp.CommandInspectorDisable:
		return emptyVal, nil

	case cdp.EventInspectorDetached:
		v = new(inspector.EventDetached)

	case cdp.EventInspectorTargetCrashed:
		v = new(inspector.EventTargetCrashed)

	case cdp.CommandMemoryGetDOMCounters:
		v = new(memory.GetDOMCountersReturns)

	case cdp.CommandMemoryPrepareForLeakDetection:
		return emptyVal, nil

	case cdp.CommandMemorySetPressureNotificationsSuppressed:
		return emptyVal, nil

	case cdp.CommandMemorySimulatePressureNotification:
		return emptyVal, nil

	case cdp.CommandPerformanceEnable:
		return emptyVal, nil

	case cdp.CommandPerformanceDisable:
		return emptyVal, nil

	case cdp.CommandPerformanceGetMetrics:
		v = new(performance.GetMetricsReturns)

	case cdp.EventPerformanceMetrics:
		v = new(performance.EventMetrics)

	case cdp.CommandPageEnable:
		return emptyVal, nil

	case cdp.CommandPageDisable:
		return emptyVal, nil

	case cdp.CommandPageAddScriptToEvaluateOnNewDocument:
		v = new(page.AddScriptToEvaluateOnNewDocumentReturns)

	case cdp.CommandPageRemoveScriptToEvaluateOnNewDocument:
		return emptyVal, nil

	case cdp.CommandPageSetAutoAttachToCreatedPages:
		return emptyVal, nil

	case cdp.CommandPageSetLifecycleEventsEnabled:
		return emptyVal, nil

	case cdp.CommandPageReload:
		return emptyVal, nil

	case cdp.CommandPageSetAdBlockingEnabled:
		return emptyVal, nil

	case cdp.CommandPageNavigate:
		v = new(page.NavigateReturns)

	case cdp.CommandPageStopLoading:
		return emptyVal, nil

	case cdp.CommandPageGetNavigationHistory:
		v = new(page.GetNavigationHistoryReturns)

	case cdp.CommandPageNavigateToHistoryEntry:
		return emptyVal, nil

	case cdp.CommandPageGetResourceTree:
		v = new(page.GetResourceTreeReturns)

	case cdp.CommandPageGetFrameTree:
		v = new(page.GetFrameTreeReturns)

	case cdp.CommandPageGetResourceContent:
		v = new(page.GetResourceContentReturns)

	case cdp.CommandPageSearchInResource:
		v = new(page.SearchInResourceReturns)

	case cdp.CommandPageSetDocumentContent:
		return emptyVal, nil

	case cdp.CommandPageCaptureScreenshot:
		v = new(page.CaptureScreenshotReturns)

	case cdp.CommandPagePrintToPDF:
		v = new(page.PrintToPDFReturns)

	case cdp.CommandPageStartScreencast:
		return emptyVal, nil

	case cdp.CommandPageStopScreencast:
		return emptyVal, nil

	case cdp.CommandPageScreencastFrameAck:
		return emptyVal, nil

	case cdp.CommandPageHandleJavaScriptDialog:
		return emptyVal, nil

	case cdp.CommandPageGetAppManifest:
		v = new(page.GetAppManifestReturns)

	case cdp.CommandPageRequestAppBanner:
		return emptyVal, nil

	case cdp.CommandPageGetLayoutMetrics:
		v = new(page.GetLayoutMetricsReturns)

	case cdp.CommandPageCreateIsolatedWorld:
		v = new(page.CreateIsolatedWorldReturns)

	case cdp.CommandPageBringToFront:
		return emptyVal, nil

	case cdp.CommandPageSetDownloadBehavior:
		return emptyVal, nil

	case cdp.EventPageDomContentEventFired:
		v = new(page.EventDomContentEventFired)

	case cdp.EventPageLoadEventFired:
		v = new(page.EventLoadEventFired)

	case cdp.EventPageLifecycleEvent:
		v = new(page.EventLifecycleEvent)

	case cdp.EventPageFrameAttached:
		v = new(page.EventFrameAttached)

	case cdp.EventPageFrameNavigated:
		v = new(page.EventFrameNavigated)

	case cdp.EventPageFrameDetached:
		v = new(page.EventFrameDetached)

	case cdp.EventPageFrameStartedLoading:
		v = new(page.EventFrameStartedLoading)

	case cdp.EventPageFrameStoppedLoading:
		v = new(page.EventFrameStoppedLoading)

	case cdp.EventPageFrameScheduledNavigation:
		v = new(page.EventFrameScheduledNavigation)

	case cdp.EventPageFrameClearedScheduledNavigation:
		v = new(page.EventFrameClearedScheduledNavigation)

	case cdp.EventPageFrameResized:
		v = new(page.EventFrameResized)

	case cdp.EventPageJavascriptDialogOpening:
		v = new(page.EventJavascriptDialogOpening)

	case cdp.EventPageJavascriptDialogClosed:
		v = new(page.EventJavascriptDialogClosed)

	case cdp.EventPageScreencastFrame:
		v = new(page.EventScreencastFrame)

	case cdp.EventPageScreencastVisibilityChanged:
		v = new(page.EventScreencastVisibilityChanged)

	case cdp.EventPageInterstitialShown:
		v = new(page.EventInterstitialShown)

	case cdp.EventPageInterstitialHidden:
		v = new(page.EventInterstitialHidden)

	case cdp.EventPageWindowOpen:
		v = new(page.EventWindowOpen)

	case cdp.CommandOverlayEnable:
		return emptyVal, nil

	case cdp.CommandOverlayDisable:
		return emptyVal, nil

	case cdp.CommandOverlaySetShowPaintRects:
		return emptyVal, nil

	case cdp.CommandOverlaySetShowDebugBorders:
		return emptyVal, nil

	case cdp.CommandOverlaySetShowFPSCounter:
		return emptyVal, nil

	case cdp.CommandOverlaySetShowScrollBottleneckRects:
		return emptyVal, nil

	case cdp.CommandOverlaySetShowViewportSizeOnResize:
		return emptyVal, nil

	case cdp.CommandOverlaySetPausedInDebuggerMessage:
		return emptyVal, nil

	case cdp.CommandOverlaySetSuspended:
		return emptyVal, nil

	case cdp.CommandOverlaySetInspectMode:
		return emptyVal, nil

	case cdp.CommandOverlayHighlightRect:
		return emptyVal, nil

	case cdp.CommandOverlayHighlightQuad:
		return emptyVal, nil

	case cdp.CommandOverlayHighlightNode:
		return emptyVal, nil

	case cdp.CommandOverlayHighlightFrame:
		return emptyVal, nil

	case cdp.CommandOverlayHideHighlight:
		return emptyVal, nil

	case cdp.CommandOverlayGetHighlightObjectForTest:
		v = new(overlay.GetHighlightObjectForTestReturns)

	case cdp.EventOverlayNodeHighlightRequested:
		v = new(overlay.EventNodeHighlightRequested)

	case cdp.EventOverlayInspectNodeRequested:
		v = new(overlay.EventInspectNodeRequested)

	case cdp.EventOverlayScreenshotRequested:
		v = new(overlay.EventScreenshotRequested)

	case cdp.CommandEmulationSetDeviceMetricsOverride:
		return emptyVal, nil

	case cdp.CommandEmulationClearDeviceMetricsOverride:
		return emptyVal, nil

	case cdp.CommandEmulationResetPageScaleFactor:
		return emptyVal, nil

	case cdp.CommandEmulationSetPageScaleFactor:
		return emptyVal, nil

	case cdp.CommandEmulationSetScriptExecutionDisabled:
		return emptyVal, nil

	case cdp.CommandEmulationSetGeolocationOverride:
		return emptyVal, nil

	case cdp.CommandEmulationClearGeolocationOverride:
		return emptyVal, nil

	case cdp.CommandEmulationSetTouchEmulationEnabled:
		return emptyVal, nil

	case cdp.CommandEmulationSetEmitTouchEventsForMouse:
		return emptyVal, nil

	case cdp.CommandEmulationSetEmulatedMedia:
		return emptyVal, nil

	case cdp.CommandEmulationSetCPUThrottlingRate:
		return emptyVal, nil

	case cdp.CommandEmulationCanEmulate:
		v = new(emulation.CanEmulateReturns)

	case cdp.CommandEmulationSetVirtualTimePolicy:
		v = new(emulation.SetVirtualTimePolicyReturns)

	case cdp.CommandEmulationSetNavigatorOverrides:
		return emptyVal, nil

	case cdp.CommandEmulationSetDefaultBackgroundColorOverride:
		return emptyVal, nil

	case cdp.EventEmulationVirtualTimeBudgetExpired:
		v = new(emulation.EventVirtualTimeBudgetExpired)

	case cdp.EventEmulationVirtualTimeAdvanced:
		v = new(emulation.EventVirtualTimeAdvanced)

	case cdp.EventEmulationVirtualTimePaused:
		v = new(emulation.EventVirtualTimePaused)

	case cdp.CommandSecurityEnable:
		return emptyVal, nil

	case cdp.CommandSecurityDisable:
		return emptyVal, nil

	case cdp.CommandSecurityHandleCertificateError:
		return emptyVal, nil

	case cdp.CommandSecuritySetOverrideCertificateErrors:
		return emptyVal, nil

	case cdp.EventSecuritySecurityStateChanged:
		v = new(security.EventSecurityStateChanged)

	case cdp.EventSecurityCertificateError:
		v = new(security.EventCertificateError)

	case cdp.CommandAuditsGetEncodedResponse:
		v = new(audits.GetEncodedResponseReturns)

	case cdp.CommandNetworkEnable:
		return emptyVal, nil

	case cdp.CommandNetworkDisable:
		return emptyVal, nil

	case cdp.CommandNetworkSetUserAgentOverride:
		return emptyVal, nil

	case cdp.CommandNetworkSearchInResponseBody:
		v = new(network.SearchInResponseBodyReturns)

	case cdp.CommandNetworkSetExtraHTTPHeaders:
		return emptyVal, nil

	case cdp.CommandNetworkGetResponseBody:
		v = new(network.GetResponseBodyReturns)

	case cdp.CommandNetworkSetBlockedURLS:
		return emptyVal, nil

	case cdp.CommandNetworkReplayXHR:
		return emptyVal, nil

	case cdp.CommandNetworkClearBrowserCache:
		return emptyVal, nil

	case cdp.CommandNetworkClearBrowserCookies:
		return emptyVal, nil

	case cdp.CommandNetworkGetCookies:
		v = new(network.GetCookiesReturns)

	case cdp.CommandNetworkGetAllCookies:
		v = new(network.GetAllCookiesReturns)

	case cdp.CommandNetworkDeleteCookies:
		return emptyVal, nil

	case cdp.CommandNetworkSetCookie:
		v = new(network.SetCookieReturns)

	case cdp.CommandNetworkSetCookies:
		return emptyVal, nil

	case cdp.CommandNetworkEmulateNetworkConditions:
		return emptyVal, nil

	case cdp.CommandNetworkSetCacheDisabled:
		return emptyVal, nil

	case cdp.CommandNetworkSetBypassServiceWorker:
		return emptyVal, nil

	case cdp.CommandNetworkSetDataSizeLimitsForTest:
		return emptyVal, nil

	case cdp.CommandNetworkGetCertificate:
		v = new(network.GetCertificateReturns)

	case cdp.CommandNetworkSetRequestInterception:
		return emptyVal, nil

	case cdp.CommandNetworkContinueInterceptedRequest:
		return emptyVal, nil

	case cdp.CommandNetworkGetResponseBodyForInterception:
		v = new(network.GetResponseBodyForInterceptionReturns)

	case cdp.EventNetworkResourceChangedPriority:
		v = new(network.EventResourceChangedPriority)

	case cdp.EventNetworkRequestWillBeSent:
		v = new(network.EventRequestWillBeSent)

	case cdp.EventNetworkRequestServedFromCache:
		v = new(network.EventRequestServedFromCache)

	case cdp.EventNetworkResponseReceived:
		v = new(network.EventResponseReceived)

	case cdp.EventNetworkDataReceived:
		v = new(network.EventDataReceived)

	case cdp.EventNetworkLoadingFinished:
		v = new(network.EventLoadingFinished)

	case cdp.EventNetworkLoadingFailed:
		v = new(network.EventLoadingFailed)

	case cdp.EventNetworkWebSocketWillSendHandshakeRequest:
		v = new(network.EventWebSocketWillSendHandshakeRequest)

	case cdp.EventNetworkWebSocketHandshakeResponseReceived:
		v = new(network.EventWebSocketHandshakeResponseReceived)

	case cdp.EventNetworkWebSocketCreated:
		v = new(network.EventWebSocketCreated)

	case cdp.EventNetworkWebSocketClosed:
		v = new(network.EventWebSocketClosed)

	case cdp.EventNetworkWebSocketFrameReceived:
		v = new(network.EventWebSocketFrameReceived)

	case cdp.EventNetworkWebSocketFrameError:
		v = new(network.EventWebSocketFrameError)

	case cdp.EventNetworkWebSocketFrameSent:
		v = new(network.EventWebSocketFrameSent)

	case cdp.EventNetworkEventSourceMessageReceived:
		v = new(network.EventEventSourceMessageReceived)

	case cdp.EventNetworkRequestIntercepted:
		v = new(network.EventRequestIntercepted)

	case cdp.CommandDatabaseEnable:
		return emptyVal, nil

	case cdp.CommandDatabaseDisable:
		return emptyVal, nil

	case cdp.CommandDatabaseGetDatabaseTableNames:
		v = new(database.GetDatabaseTableNamesReturns)

	case cdp.CommandDatabaseExecuteSQL:
		v = new(database.ExecuteSQLReturns)

	case cdp.EventDatabaseAddDatabase:
		v = new(database.EventAddDatabase)

	case cdp.CommandIndexedDBEnable:
		return emptyVal, nil

	case cdp.CommandIndexedDBDisable:
		return emptyVal, nil

	case cdp.CommandIndexedDBRequestDatabaseNames:
		v = new(indexeddb.RequestDatabaseNamesReturns)

	case cdp.CommandIndexedDBRequestDatabase:
		v = new(indexeddb.RequestDatabaseReturns)

	case cdp.CommandIndexedDBRequestData:
		v = new(indexeddb.RequestDataReturns)

	case cdp.CommandIndexedDBClearObjectStore:
		return emptyVal, nil

	case cdp.CommandIndexedDBDeleteDatabase:
		return emptyVal, nil

	case cdp.CommandCacheStorageRequestCacheNames:
		v = new(cachestorage.RequestCacheNamesReturns)

	case cdp.CommandCacheStorageRequestEntries:
		v = new(cachestorage.RequestEntriesReturns)

	case cdp.CommandCacheStorageDeleteCache:
		return emptyVal, nil

	case cdp.CommandCacheStorageDeleteEntry:
		return emptyVal, nil

	case cdp.CommandCacheStorageRequestCachedResponse:
		v = new(cachestorage.RequestCachedResponseReturns)

	case cdp.CommandDOMStorageEnable:
		return emptyVal, nil

	case cdp.CommandDOMStorageDisable:
		return emptyVal, nil

	case cdp.CommandDOMStorageClear:
		return emptyVal, nil

	case cdp.CommandDOMStorageGetDOMStorageItems:
		v = new(domstorage.GetDOMStorageItemsReturns)

	case cdp.CommandDOMStorageSetDOMStorageItem:
		return emptyVal, nil

	case cdp.CommandDOMStorageRemoveDOMStorageItem:
		return emptyVal, nil

	case cdp.EventDOMStorageDomStorageItemsCleared:
		v = new(domstorage.EventDomStorageItemsCleared)

	case cdp.EventDOMStorageDomStorageItemRemoved:
		v = new(domstorage.EventDomStorageItemRemoved)

	case cdp.EventDOMStorageDomStorageItemAdded:
		v = new(domstorage.EventDomStorageItemAdded)

	case cdp.EventDOMStorageDomStorageItemUpdated:
		v = new(domstorage.EventDomStorageItemUpdated)

	case cdp.CommandApplicationCacheGetFramesWithManifests:
		v = new(applicationcache.GetFramesWithManifestsReturns)

	case cdp.CommandApplicationCacheEnable:
		return emptyVal, nil

	case cdp.CommandApplicationCacheGetManifestForFrame:
		v = new(applicationcache.GetManifestForFrameReturns)

	case cdp.CommandApplicationCacheGetApplicationCacheForFrame:
		v = new(applicationcache.GetApplicationCacheForFrameReturns)

	case cdp.EventApplicationCacheApplicationCacheStatusUpdated:
		v = new(applicationcache.EventApplicationCacheStatusUpdated)

	case cdp.EventApplicationCacheNetworkStateUpdated:
		v = new(applicationcache.EventNetworkStateUpdated)

	case cdp.CommandDOMEnable:
		return emptyVal, nil

	case cdp.CommandDOMDisable:
		return emptyVal, nil

	case cdp.CommandDOMGetDocument:
		v = new(dom.GetDocumentReturns)

	case cdp.CommandDOMGetFlattenedDocument:
		v = new(dom.GetFlattenedDocumentReturns)

	case cdp.CommandDOMCollectClassNamesFromSubtree:
		v = new(dom.CollectClassNamesFromSubtreeReturns)

	case cdp.CommandDOMRequestChildNodes:
		return emptyVal, nil

	case cdp.CommandDOMQuerySelector:
		v = new(dom.QuerySelectorReturns)

	case cdp.CommandDOMQuerySelectorAll:
		v = new(dom.QuerySelectorAllReturns)

	case cdp.CommandDOMSetNodeName:
		v = new(dom.SetNodeNameReturns)

	case cdp.CommandDOMSetNodeValue:
		return emptyVal, nil

	case cdp.CommandDOMRemoveNode:
		return emptyVal, nil

	case cdp.CommandDOMSetAttributeValue:
		return emptyVal, nil

	case cdp.CommandDOMSetAttributesAsText:
		return emptyVal, nil

	case cdp.CommandDOMRemoveAttribute:
		return emptyVal, nil

	case cdp.CommandDOMGetOuterHTML:
		v = new(dom.GetOuterHTMLReturns)

	case cdp.CommandDOMSetOuterHTML:
		return emptyVal, nil

	case cdp.CommandDOMPerformSearch:
		v = new(dom.PerformSearchReturns)

	case cdp.CommandDOMGetSearchResults:
		v = new(dom.GetSearchResultsReturns)

	case cdp.CommandDOMDiscardSearchResults:
		return emptyVal, nil

	case cdp.CommandDOMRequestNode:
		v = new(dom.RequestNodeReturns)

	case cdp.CommandDOMPushNodeByPathToFrontend:
		v = new(dom.PushNodeByPathToFrontendReturns)

	case cdp.CommandDOMPushNodesByBackendIdsToFrontend:
		v = new(dom.PushNodesByBackendIdsToFrontendReturns)

	case cdp.CommandDOMSetInspectedNode:
		return emptyVal, nil

	case cdp.CommandDOMResolveNode:
		v = new(dom.ResolveNodeReturns)

	case cdp.CommandDOMGetAttributes:
		v = new(dom.GetAttributesReturns)

	case cdp.CommandDOMCopyTo:
		v = new(dom.CopyToReturns)

	case cdp.CommandDOMMoveTo:
		v = new(dom.MoveToReturns)

	case cdp.CommandDOMUndo:
		return emptyVal, nil

	case cdp.CommandDOMRedo:
		return emptyVal, nil

	case cdp.CommandDOMMarkUndoableState:
		return emptyVal, nil

	case cdp.CommandDOMFocus:
		return emptyVal, nil

	case cdp.CommandDOMSetFileInputFiles:
		return emptyVal, nil

	case cdp.CommandDOMGetBoxModel:
		v = new(dom.GetBoxModelReturns)

	case cdp.CommandDOMGetNodeForLocation:
		v = new(dom.GetNodeForLocationReturns)

	case cdp.CommandDOMGetRelayoutBoundary:
		v = new(dom.GetRelayoutBoundaryReturns)

	case cdp.CommandDOMDescribeNode:
		v = new(dom.DescribeNodeReturns)

	case cdp.EventDOMDocumentUpdated:
		v = new(dom.EventDocumentUpdated)

	case cdp.EventDOMSetChildNodes:
		v = new(dom.EventSetChildNodes)

	case cdp.EventDOMAttributeModified:
		v = new(dom.EventAttributeModified)

	case cdp.EventDOMAttributeRemoved:
		v = new(dom.EventAttributeRemoved)

	case cdp.EventDOMInlineStyleInvalidated:
		v = new(dom.EventInlineStyleInvalidated)

	case cdp.EventDOMCharacterDataModified:
		v = new(dom.EventCharacterDataModified)

	case cdp.EventDOMChildNodeCountUpdated:
		v = new(dom.EventChildNodeCountUpdated)

	case cdp.EventDOMChildNodeInserted:
		v = new(dom.EventChildNodeInserted)

	case cdp.EventDOMChildNodeRemoved:
		v = new(dom.EventChildNodeRemoved)

	case cdp.EventDOMShadowRootPushed:
		v = new(dom.EventShadowRootPushed)

	case cdp.EventDOMShadowRootPopped:
		v = new(dom.EventShadowRootPopped)

	case cdp.EventDOMPseudoElementAdded:
		v = new(dom.EventPseudoElementAdded)

	case cdp.EventDOMPseudoElementRemoved:
		v = new(dom.EventPseudoElementRemoved)

	case cdp.EventDOMDistributedNodesUpdated:
		v = new(dom.EventDistributedNodesUpdated)

	case cdp.CommandCSSEnable:
		return emptyVal, nil

	case cdp.CommandCSSDisable:
		return emptyVal, nil

	case cdp.CommandCSSGetMatchedStylesForNode:
		v = new(css.GetMatchedStylesForNodeReturns)

	case cdp.CommandCSSGetInlineStylesForNode:
		v = new(css.GetInlineStylesForNodeReturns)

	case cdp.CommandCSSGetComputedStyleForNode:
		v = new(css.GetComputedStyleForNodeReturns)

	case cdp.CommandCSSGetPlatformFontsForNode:
		v = new(css.GetPlatformFontsForNodeReturns)

	case cdp.CommandCSSGetStyleSheetText:
		v = new(css.GetStyleSheetTextReturns)

	case cdp.CommandCSSCollectClassNames:
		v = new(css.CollectClassNamesReturns)

	case cdp.CommandCSSSetStyleSheetText:
		v = new(css.SetStyleSheetTextReturns)

	case cdp.CommandCSSSetRuleSelector:
		v = new(css.SetRuleSelectorReturns)

	case cdp.CommandCSSSetKeyframeKey:
		v = new(css.SetKeyframeKeyReturns)

	case cdp.CommandCSSSetStyleTexts:
		v = new(css.SetStyleTextsReturns)

	case cdp.CommandCSSSetMediaText:
		v = new(css.SetMediaTextReturns)

	case cdp.CommandCSSCreateStyleSheet:
		v = new(css.CreateStyleSheetReturns)

	case cdp.CommandCSSAddRule:
		v = new(css.AddRuleReturns)

	case cdp.CommandCSSForcePseudoState:
		return emptyVal, nil

	case cdp.CommandCSSGetMediaQueries:
		v = new(css.GetMediaQueriesReturns)

	case cdp.CommandCSSSetEffectivePropertyValueForNode:
		return emptyVal, nil

	case cdp.CommandCSSGetBackgroundColors:
		v = new(css.GetBackgroundColorsReturns)

	case cdp.CommandCSSStartRuleUsageTracking:
		return emptyVal, nil

	case cdp.CommandCSSTakeCoverageDelta:
		v = new(css.TakeCoverageDeltaReturns)

	case cdp.CommandCSSStopRuleUsageTracking:
		v = new(css.StopRuleUsageTrackingReturns)

	case cdp.EventCSSMediaQueryResultChanged:
		v = new(css.EventMediaQueryResultChanged)

	case cdp.EventCSSFontsUpdated:
		v = new(css.EventFontsUpdated)

	case cdp.EventCSSStyleSheetChanged:
		v = new(css.EventStyleSheetChanged)

	case cdp.EventCSSStyleSheetAdded:
		v = new(css.EventStyleSheetAdded)

	case cdp.EventCSSStyleSheetRemoved:
		v = new(css.EventStyleSheetRemoved)

	case cdp.CommandDOMSnapshotGetSnapshot:
		v = new(domsnapshot.GetSnapshotReturns)

	case cdp.CommandIORead:
		v = new(iodom.ReadReturns)

	case cdp.CommandIOClose:
		return emptyVal, nil

	case cdp.CommandIOResolveBlob:
		v = new(iodom.ResolveBlobReturns)

	case cdp.CommandDOMDebuggerSetDOMBreakpoint:
		return emptyVal, nil

	case cdp.CommandDOMDebuggerRemoveDOMBreakpoint:
		return emptyVal, nil

	case cdp.CommandDOMDebuggerSetEventListenerBreakpoint:
		return emptyVal, nil

	case cdp.CommandDOMDebuggerRemoveEventListenerBreakpoint:
		return emptyVal, nil

	case cdp.CommandDOMDebuggerSetInstrumentationBreakpoint:
		return emptyVal, nil

	case cdp.CommandDOMDebuggerRemoveInstrumentationBreakpoint:
		return emptyVal, nil

	case cdp.CommandDOMDebuggerSetXHRBreakpoint:
		return emptyVal, nil

	case cdp.CommandDOMDebuggerRemoveXHRBreakpoint:
		return emptyVal, nil

	case cdp.CommandDOMDebuggerGetEventListeners:
		v = new(domdebugger.GetEventListenersReturns)

	case cdp.CommandTargetSetDiscoverTargets:
		return emptyVal, nil

	case cdp.CommandTargetSetAutoAttach:
		return emptyVal, nil

	case cdp.CommandTargetSetAttachToFrames:
		return emptyVal, nil

	case cdp.CommandTargetSetRemoteLocations:
		return emptyVal, nil

	case cdp.CommandTargetSendMessageToTarget:
		return emptyVal, nil

	case cdp.CommandTargetGetTargetInfo:
		v = new(target.GetTargetInfoReturns)

	case cdp.CommandTargetActivateTarget:
		return emptyVal, nil

	case cdp.CommandTargetCloseTarget:
		v = new(target.CloseTargetReturns)

	case cdp.CommandTargetAttachToTarget:
		v = new(target.AttachToTargetReturns)

	case cdp.CommandTargetDetachFromTarget:
		return emptyVal, nil

	case cdp.CommandTargetCreateBrowserContext:
		v = new(target.CreateBrowserContextReturns)

	case cdp.CommandTargetDisposeBrowserContext:
		v = new(target.DisposeBrowserContextReturns)

	case cdp.CommandTargetCreateTarget:
		v = new(target.CreateTargetReturns)

	case cdp.CommandTargetGetTargets:
		v = new(target.GetTargetsReturns)

	case cdp.EventTargetTargetCreated:
		v = new(target.EventTargetCreated)

	case cdp.EventTargetTargetInfoChanged:
		v = new(target.EventTargetInfoChanged)

	case cdp.EventTargetTargetDestroyed:
		v = new(target.EventTargetDestroyed)

	case cdp.EventTargetAttachedToTarget:
		v = new(target.EventAttachedToTarget)

	case cdp.EventTargetDetachedFromTarget:
		v = new(target.EventDetachedFromTarget)

	case cdp.EventTargetReceivedMessageFromTarget:
		v = new(target.EventReceivedMessageFromTarget)

	case cdp.CommandHeadlessExperimentalEnable:
		return emptyVal, nil

	case cdp.CommandHeadlessExperimentalDisable:
		return emptyVal, nil

	case cdp.CommandHeadlessExperimentalBeginFrame:
		v = new(headlessexperimental.BeginFrameReturns)

	case cdp.EventHeadlessExperimentalNeedsBeginFramesChanged:
		v = new(headlessexperimental.EventNeedsBeginFramesChanged)

	case cdp.EventHeadlessExperimentalMainFrameReadyForScreenshots:
		v = new(headlessexperimental.EventMainFrameReadyForScreenshots)

	case cdp.CommandServiceWorkerEnable:
		return emptyVal, nil

	case cdp.CommandServiceWorkerDisable:
		return emptyVal, nil

	case cdp.CommandServiceWorkerUnregister:
		return emptyVal, nil

	case cdp.CommandServiceWorkerUpdateRegistration:
		return emptyVal, nil

	case cdp.CommandServiceWorkerStartWorker:
		return emptyVal, nil

	case cdp.CommandServiceWorkerSkipWaiting:
		return emptyVal, nil

	case cdp.CommandServiceWorkerStopWorker:
		return emptyVal, nil

	case cdp.CommandServiceWorkerStopAllWorkers:
		return emptyVal, nil

	case cdp.CommandServiceWorkerInspectWorker:
		return emptyVal, nil

	case cdp.CommandServiceWorkerSetForceUpdateOnPageLoad:
		return emptyVal, nil

	case cdp.CommandServiceWorkerDeliverPushMessage:
		return emptyVal, nil

	case cdp.CommandServiceWorkerDispatchSyncEvent:
		return emptyVal, nil

	case cdp.EventServiceWorkerWorkerRegistrationUpdated:
		v = new(serviceworker.EventWorkerRegistrationUpdated)

	case cdp.EventServiceWorkerWorkerVersionUpdated:
		v = new(serviceworker.EventWorkerVersionUpdated)

	case cdp.EventServiceWorkerWorkerErrorReported:
		v = new(serviceworker.EventWorkerErrorReported)

	case cdp.CommandInputSetIgnoreInputEvents:
		return emptyVal, nil

	case cdp.CommandInputDispatchKeyEvent:
		return emptyVal, nil

	case cdp.CommandInputDispatchMouseEvent:
		return emptyVal, nil

	case cdp.CommandInputDispatchTouchEvent:
		return emptyVal, nil

	case cdp.CommandInputEmulateTouchFromMouseEvent:
		return emptyVal, nil

	case cdp.CommandInputSynthesizePinchGesture:
		return emptyVal, nil

	case cdp.CommandInputSynthesizeScrollGesture:
		return emptyVal, nil

	case cdp.CommandInputSynthesizeTapGesture:
		return emptyVal, nil

	case cdp.CommandLayerTreeEnable:
		return emptyVal, nil

	case cdp.CommandLayerTreeDisable:
		return emptyVal, nil

	case cdp.CommandLayerTreeCompositingReasons:
		v = new(layertree.CompositingReasonsReturns)

	case cdp.CommandLayerTreeMakeSnapshot:
		v = new(layertree.MakeSnapshotReturns)

	case cdp.CommandLayerTreeLoadSnapshot:
		v = new(layertree.LoadSnapshotReturns)

	case cdp.CommandLayerTreeReleaseSnapshot:
		return emptyVal, nil

	case cdp.CommandLayerTreeProfileSnapshot:
		v = new(layertree.ProfileSnapshotReturns)

	case cdp.CommandLayerTreeReplaySnapshot:
		v = new(layertree.ReplaySnapshotReturns)

	case cdp.CommandLayerTreeSnapshotCommandLog:
		v = new(layertree.SnapshotCommandLogReturns)

	case cdp.EventLayerTreeLayerTreeDidChange:
		v = new(layertree.EventLayerTreeDidChange)

	case cdp.EventLayerTreeLayerPainted:
		v = new(layertree.EventLayerPainted)

	case cdp.CommandDeviceOrientationSetDeviceOrientationOverride:
		return emptyVal, nil

	case cdp.CommandDeviceOrientationClearDeviceOrientationOverride:
		return emptyVal, nil

	case cdp.CommandTracingStart:
		return emptyVal, nil

	case cdp.CommandTracingEnd:
		return emptyVal, nil

	case cdp.CommandTracingGetCategories:
		v = new(tracing.GetCategoriesReturns)

	case cdp.CommandTracingRequestMemoryDump:
		v = new(tracing.RequestMemoryDumpReturns)

	case cdp.CommandTracingRecordClockSyncMarker:
		return emptyVal, nil

	case cdp.EventTracingDataCollected:
		v = new(tracing.EventDataCollected)

	case cdp.EventTracingTracingComplete:
		v = new(tracing.EventTracingComplete)

	case cdp.EventTracingBufferUsage:
		v = new(tracing.EventBufferUsage)

	case cdp.CommandAnimationEnable:
		return emptyVal, nil

	case cdp.CommandAnimationDisable:
		return emptyVal, nil

	case cdp.CommandAnimationGetPlaybackRate:
		v = new(animation.GetPlaybackRateReturns)

	case cdp.CommandAnimationSetPlaybackRate:
		return emptyVal, nil

	case cdp.CommandAnimationGetCurrentTime:
		v = new(animation.GetCurrentTimeReturns)

	case cdp.CommandAnimationSetPaused:
		return emptyVal, nil

	case cdp.CommandAnimationSetTiming:
		return emptyVal, nil

	case cdp.CommandAnimationSeekAnimations:
		return emptyVal, nil

	case cdp.CommandAnimationReleaseAnimations:
		return emptyVal, nil

	case cdp.CommandAnimationResolveAnimation:
		v = new(animation.ResolveAnimationReturns)

	case cdp.EventAnimationAnimationCreated:
		v = new(animation.EventAnimationCreated)

	case cdp.EventAnimationAnimationStarted:
		v = new(animation.EventAnimationStarted)

	case cdp.EventAnimationAnimationCanceled:
		v = new(animation.EventAnimationCanceled)

	case cdp.CommandAccessibilityGetPartialAXTree:
		v = new(accessibility.GetPartialAXTreeReturns)

	case cdp.CommandStorageClearDataForOrigin:
		return emptyVal, nil

	case cdp.CommandStorageGetUsageAndQuota:
		v = new(storage.GetUsageAndQuotaReturns)

	case cdp.CommandStorageTrackCacheStorageForOrigin:
		return emptyVal, nil

	case cdp.CommandStorageUntrackCacheStorageForOrigin:
		return emptyVal, nil

	case cdp.CommandStorageTrackIndexedDBForOrigin:
		return emptyVal, nil

	case cdp.CommandStorageUntrackIndexedDBForOrigin:
		return emptyVal, nil

	case cdp.EventStorageCacheStorageListUpdated:
		v = new(storage.EventCacheStorageListUpdated)

	case cdp.EventStorageCacheStorageContentUpdated:
		v = new(storage.EventCacheStorageContentUpdated)

	case cdp.EventStorageIndexedDBListUpdated:
		v = new(storage.EventIndexedDBListUpdated)

	case cdp.EventStorageIndexedDBContentUpdated:
		v = new(storage.EventIndexedDBContentUpdated)

	case cdp.CommandLogEnable:
		return emptyVal, nil

	case cdp.CommandLogDisable:
		return emptyVal, nil

	case cdp.CommandLogClear:
		return emptyVal, nil

	case cdp.CommandLogStartViolationsReport:
		return emptyVal, nil

	case cdp.CommandLogStopViolationsReport:
		return emptyVal, nil

	case cdp.EventLogEntryAdded:
		v = new(logdom.EventEntryAdded)

	case cdp.CommandSystemInfoGetInfo:
		v = new(systeminfo.GetInfoReturns)

	case cdp.CommandTetheringBind:
		return emptyVal, nil

	case cdp.CommandTetheringUnbind:
		return emptyVal, nil

	case cdp.EventTetheringAccepted:
		v = new(tethering.EventAccepted)

	case cdp.CommandBrowserClose:
		return emptyVal, nil

	case cdp.CommandBrowserGetWindowForTarget:
		v = new(browser.GetWindowForTargetReturns)

	case cdp.CommandBrowserGetVersion:
		v = new(browser.GetVersionReturns)

	case cdp.CommandBrowserSetWindowBounds:
		return emptyVal, nil

	case cdp.CommandBrowserGetWindowBounds:
		v = new(browser.GetWindowBoundsReturns)

	case cdp.CommandSchemaGetDomains:
		v = new(schema.GetDomainsReturns)

	case cdp.CommandRuntimeEvaluate:
		v = new(runtime.EvaluateReturns)

	case cdp.CommandRuntimeAwaitPromise:
		v = new(runtime.AwaitPromiseReturns)

	case cdp.CommandRuntimeCallFunctionOn:
		v = new(runtime.CallFunctionOnReturns)

	case cdp.CommandRuntimeGetProperties:
		v = new(runtime.GetPropertiesReturns)

	case cdp.CommandRuntimeReleaseObject:
		return emptyVal, nil

	case cdp.CommandRuntimeReleaseObjectGroup:
		return emptyVal, nil

	case cdp.CommandRuntimeRunIfWaitingForDebugger:
		return emptyVal, nil

	case cdp.CommandRuntimeEnable:
		return emptyVal, nil

	case cdp.CommandRuntimeDisable:
		return emptyVal, nil

	case cdp.CommandRuntimeDiscardConsoleEntries:
		return emptyVal, nil

	case cdp.CommandRuntimeSetCustomObjectFormatterEnabled:
		return emptyVal, nil

	case cdp.CommandRuntimeCompileScript:
		v = new(runtime.CompileScriptReturns)

	case cdp.CommandRuntimeRunScript:
		v = new(runtime.RunScriptReturns)

	case cdp.CommandRuntimeQueryObjects:
		v = new(runtime.QueryObjectsReturns)

	case cdp.CommandRuntimeGlobalLexicalScopeNames:
		v = new(runtime.GlobalLexicalScopeNamesReturns)

	case cdp.EventRuntimeExecutionContextCreated:
		v = new(runtime.EventExecutionContextCreated)

	case cdp.EventRuntimeExecutionContextDestroyed:
		v = new(runtime.EventExecutionContextDestroyed)

	case cdp.EventRuntimeExecutionContextsCleared:
		v = new(runtime.EventExecutionContextsCleared)

	case cdp.EventRuntimeExceptionThrown:
		v = new(runtime.EventExceptionThrown)

	case cdp.EventRuntimeExceptionRevoked:
		v = new(runtime.EventExceptionRevoked)

	case cdp.EventRuntimeConsoleAPICalled:
		v = new(runtime.EventConsoleAPICalled)

	case cdp.EventRuntimeInspectRequested:
		v = new(runtime.EventInspectRequested)

	case cdp.CommandDebuggerEnable:
		return emptyVal, nil

	case cdp.CommandDebuggerDisable:
		return emptyVal, nil

	case cdp.CommandDebuggerSetBreakpointsActive:
		return emptyVal, nil

	case cdp.CommandDebuggerSetSkipAllPauses:
		return emptyVal, nil

	case cdp.CommandDebuggerSetBreakpointByURL:
		v = new(debugger.SetBreakpointByURLReturns)

	case cdp.CommandDebuggerSetBreakpoint:
		v = new(debugger.SetBreakpointReturns)

	case cdp.CommandDebuggerRemoveBreakpoint:
		return emptyVal, nil

	case cdp.CommandDebuggerGetPossibleBreakpoints:
		v = new(debugger.GetPossibleBreakpointsReturns)

	case cdp.CommandDebuggerContinueToLocation:
		return emptyVal, nil

	case cdp.CommandDebuggerPauseOnAsyncTask:
		return emptyVal, nil

	case cdp.CommandDebuggerStepOver:
		return emptyVal, nil

	case cdp.CommandDebuggerStepInto:
		return emptyVal, nil

	case cdp.CommandDebuggerStepOut:
		return emptyVal, nil

	case cdp.CommandDebuggerPause:
		return emptyVal, nil

	case cdp.CommandDebuggerScheduleStepIntoAsync:
		return emptyVal, nil

	case cdp.CommandDebuggerResume:
		return emptyVal, nil

	case cdp.CommandDebuggerSearchInContent:
		v = new(debugger.SearchInContentReturns)

	case cdp.CommandDebuggerSetScriptSource:
		v = new(debugger.SetScriptSourceReturns)

	case cdp.CommandDebuggerRestartFrame:
		v = new(debugger.RestartFrameReturns)

	case cdp.CommandDebuggerGetScriptSource:
		v = new(debugger.GetScriptSourceReturns)

	case cdp.CommandDebuggerSetPauseOnExceptions:
		return emptyVal, nil

	case cdp.CommandDebuggerEvaluateOnCallFrame:
		v = new(debugger.EvaluateOnCallFrameReturns)

	case cdp.CommandDebuggerSetVariableValue:
		return emptyVal, nil

	case cdp.CommandDebuggerSetReturnValue:
		return emptyVal, nil

	case cdp.CommandDebuggerSetAsyncCallStackDepth:
		return emptyVal, nil

	case cdp.CommandDebuggerSetBlackboxPatterns:
		return emptyVal, nil

	case cdp.CommandDebuggerSetBlackboxedRanges:
		return emptyVal, nil

	case cdp.EventDebuggerScriptParsed:
		v = new(debugger.EventScriptParsed)

	case cdp.EventDebuggerScriptFailedToParse:
		v = new(debugger.EventScriptFailedToParse)

	case cdp.EventDebuggerBreakpointResolved:
		v = new(debugger.EventBreakpointResolved)

	case cdp.EventDebuggerPaused:
		v = new(debugger.EventPaused)

	case cdp.EventDebuggerResumed:
		v = new(debugger.EventResumed)

	case cdp.CommandProfilerEnable:
		return emptyVal, nil

	case cdp.CommandProfilerDisable:
		return emptyVal, nil

	case cdp.CommandProfilerSetSamplingInterval:
		return emptyVal, nil

	case cdp.CommandProfilerStart:
		return emptyVal, nil

	case cdp.CommandProfilerStop:
		v = new(profiler.StopReturns)

	case cdp.CommandProfilerStartPreciseCoverage:
		return emptyVal, nil

	case cdp.CommandProfilerStopPreciseCoverage:
		return emptyVal, nil

	case cdp.CommandProfilerTakePreciseCoverage:
		v = new(profiler.TakePreciseCoverageReturns)

	case cdp.CommandProfilerGetBestEffortCoverage:
		v = new(profiler.GetBestEffortCoverageReturns)

	case cdp.CommandProfilerStartTypeProfile:
		return emptyVal, nil

	case cdp.CommandProfilerStopTypeProfile:
		return emptyVal, nil

	case cdp.CommandProfilerTakeTypeProfile:
		v = new(profiler.TakeTypeProfileReturns)

	case cdp.EventProfilerConsoleProfileStarted:
		v = new(profiler.EventConsoleProfileStarted)

	case cdp.EventProfilerConsoleProfileFinished:
		v = new(profiler.EventConsoleProfileFinished)

	case cdp.CommandHeapProfilerEnable:
		return emptyVal, nil

	case cdp.CommandHeapProfilerDisable:
		return emptyVal, nil

	case cdp.CommandHeapProfilerStartTrackingHeapObjects:
		return emptyVal, nil

	case cdp.CommandHeapProfilerStopTrackingHeapObjects:
		return emptyVal, nil

	case cdp.CommandHeapProfilerTakeHeapSnapshot:
		return emptyVal, nil

	case cdp.CommandHeapProfilerCollectGarbage:
		return emptyVal, nil

	case cdp.CommandHeapProfilerGetObjectByHeapObjectID:
		v = new(heapprofiler.GetObjectByHeapObjectIDReturns)

	case cdp.CommandHeapProfilerAddInspectedHeapObject:
		return emptyVal, nil

	case cdp.CommandHeapProfilerGetHeapObjectID:
		v = new(heapprofiler.GetHeapObjectIDReturns)

	case cdp.CommandHeapProfilerStartSampling:
		return emptyVal, nil

	case cdp.CommandHeapProfilerStopSampling:
		v = new(heapprofiler.StopSamplingReturns)

	case cdp.CommandHeapProfilerGetSamplingProfile:
		v = new(heapprofiler.GetSamplingProfileReturns)

	case cdp.EventHeapProfilerAddHeapSnapshotChunk:
		v = new(heapprofiler.EventAddHeapSnapshotChunk)

	case cdp.EventHeapProfilerResetProfiles:
		v = new(heapprofiler.EventResetProfiles)

	case cdp.EventHeapProfilerReportHeapSnapshotProgress:
		v = new(heapprofiler.EventReportHeapSnapshotProgress)

	case cdp.EventHeapProfilerLastSeenObjectID:
		v = new(heapprofiler.EventLastSeenObjectID)

	case cdp.EventHeapProfilerHeapStatsUpdate:
		v = new(heapprofiler.EventHeapStatsUpdate)
	}

	var buf easyjson.RawMessage
	switch {
	case msg.Params != nil:
		buf = msg.Params

	case msg.Result != nil:
		buf = msg.Result

	default:
		return nil, errors.New("msg missing params or result")
	}

	err := easyjson.Unmarshal(buf, v)
	if err != nil {
		return nil, err
	}

	return v, nil
}
