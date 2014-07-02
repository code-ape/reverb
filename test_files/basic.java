package org.scribe.oauth;

import java.util.*;

import org.scribe.builder.api.*;
import org.scribe.model.*;
import org.scribe.services.*;
import org.scribe.utils.*;
import java.util.concurrent.TimeUnit;

/**
 * OAuth 1.0a implementation of {@link OAuthService}
 *
 * @author Pablo Fernandez
 */
public class OAuth10aServiceImpl implements OAuthService
{
  private static final String VERSION = "1.0";

  private OAuthConfig config;
  private DefaultApi10a api;

  /**
   * Default constructor
   *
   * @param api OAuth1.0a api information
   * @param config OAuth 1.0a configuration param object
   */
  public OAuth10aServiceImpl(DefaultApi10a api, OAuthConfig config)
  {
    this.api = api;
    this.config = config; //single line comment
  }

  /**
   * {@inheritDoc}
   */
  public Token getRequestToken(int timeout, TimeUnit unit)
  {
    return getRequestToken(new TimeoutTuner(timeout, unit));
  }

  public Token getRequestToken()
  {
    return getRequestToken(2, TimeUnit.SECONDS);
  }

  public Token getRequestToken(RequestTuner tuner)
  {
    config.log("obtaining request token from " + api.getRequestTokenEndpoint());
    OAuthRequest request = new OAuthRequest(api.getRequestTokenVerb(), api.getRequestTokenEndpoint());

    config.log("setting oauth_callback to " + config.getCallback());
    request.addOAuthParameter(OAuthConstants.CALLBACK, config.getCallback());
    addOAuthParams(request, OAuthConstants.EMPTY_TOKEN);
    appendSignature(request);

    config.log("sending request...");
    Response response = request.send(tuner);
    String body = response.getBody();

    config.log("response status code: " + response.getCode());
    config.log("response body: " + body);
    return api.getRequestTokenExtractor().extract(body);
  }
}
